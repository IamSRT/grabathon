package com.grabathon.paymentassistant.storage.cache.memory;

import com.grabathon.paymentassistant.storage.db.entity.Action;
import com.grabathon.paymentassistant.storage.db.entity.Step;
import com.grabathon.paymentassistant.storage.db.entity.Template;
import com.grabathon.paymentassistant.storage.db.repository.ActionRepository;
import com.grabathon.paymentassistant.storage.db.repository.StepRepository;
import com.grabathon.paymentassistant.storage.db.repository.TemplateRepository;
import org.apache.logging.log4j.LogManager;
import org.apache.logging.log4j.Logger;
import org.springframework.beans.factory.annotation.Autowired;
import org.springframework.stereotype.Repository;
import org.springframework.util.StringUtils;

import java.util.Arrays;
import java.util.List;
import java.util.Map;
import java.util.concurrent.ConcurrentHashMap;
import java.util.concurrent.locks.Lock;
import java.util.concurrent.locks.ReentrantLock;
import java.util.stream.Collectors;


@Repository public class Cache {

    private static final Logger logger = LogManager.getLogger(Cache.class);

    private Lock lock = new ReentrantLock();

    @Autowired private StepRepository stepRepository;
    @Autowired private ActionRepository actionRepository;
    @Autowired private TemplateRepository templateRepository;


    private static Map<Long, com.grabathon.paymentassistant.storage.db.wrapper.Action> actionCache = new ConcurrentHashMap<>();
    public static Map<Long, com.grabathon.paymentassistant.storage.db.wrapper.Action> getActionCache() { return actionCache; }

    private static Map<Long, com.grabathon.paymentassistant.storage.db.wrapper.Template> templateCache = new ConcurrentHashMap<>();
    public static Map<Long, com.grabathon.paymentassistant.storage.db.wrapper.Template> getTemplateCache() { return templateCache; }

    private static Map<Long, com.grabathon.paymentassistant.storage.db.wrapper.Step> stepCache = new ConcurrentHashMap<>();
    public static Map<Long, com.grabathon.paymentassistant.storage.db.wrapper.Step> getStepCache() { return stepCache; }

    private static Map<Long, List<com.grabathon.paymentassistant.storage.db.wrapper.Step>> nextStepCache = new ConcurrentHashMap<>();
    public static Map<Long, List<com.grabathon.paymentassistant.storage.db.wrapper.Step>> getNextStepCache() { return nextStepCache; }


    public void load () {
        Map<Long, com.grabathon.paymentassistant.storage.db.wrapper.Action> actionCache = buildActionCache();
        Map<Long, com.grabathon.paymentassistant.storage.db.wrapper.Template> templateCache = buildTemplateCache();
        Map<Long, com.grabathon.paymentassistant.storage.db.wrapper.Step> stepCache = buildStepCache(actionCache, templateCache);
        Map<Long, List<com.grabathon.paymentassistant.storage.db.wrapper.Step>> nextStepCache = buildNextStepCache(stepCache);

        updateCache(
                actionCache,
                templateCache,
                stepCache,
                nextStepCache
        );
    }

    public Map<Long, com.grabathon.paymentassistant.storage.db.wrapper.Action> buildActionCache() {
        Map<Long, com.grabathon.paymentassistant.storage.db.wrapper.Action> newCache = new ConcurrentHashMap<>();
        List<Action> all = actionRepository.findAll();
        for (Action a : all) {
            newCache.put(a.getId(), new com.grabathon.paymentassistant.storage.db.wrapper.Action(a));
        }
        return newCache;
    }

    public Map<Long, com.grabathon.paymentassistant.storage.db.wrapper.Template> buildTemplateCache() {
        Map<Long, com.grabathon.paymentassistant.storage.db.wrapper.Template> newCache = new ConcurrentHashMap<>();
        List<Template> all = templateRepository.findAll();
        for (Template t : all) {
            newCache.put(t.getId(), new com.grabathon.paymentassistant.storage.db.wrapper.Template(t));
        }
        return newCache;
    }

    public Map<Long, com.grabathon.paymentassistant.storage.db.wrapper.Step> buildStepCache(
            Map<Long, com.grabathon.paymentassistant.storage.db.wrapper.Action> actionCache,
            Map<Long, com.grabathon.paymentassistant.storage.db.wrapper.Template> templateCache) {

        Map<Long, com.grabathon.paymentassistant.storage.db.wrapper.Step> newCache = new ConcurrentHashMap<>();
        List<Step> all = stepRepository.findAll();
        for (Step s : all) {
            List<com.grabathon.paymentassistant.storage.db.wrapper.Action> actions =
                    Arrays.stream(s.getActions().split(",")).filter(id -> !StringUtils.isEmpty(id)).map(id -> actionCache.get(Long.parseLong(id))).collect(Collectors.toList());

            List<com.grabathon.paymentassistant.storage.db.wrapper.Template> templates =
                    Arrays.stream(s.getTemplates().split(",")).filter(id -> !StringUtils.isEmpty(id)).map(id -> templateCache.get(Long.parseLong(id))).collect(Collectors.toList());

            newCache.put(s.getId(), new com.grabathon.paymentassistant.storage.db.wrapper.Step(
                    s.getId(),
                    s.getTitle(),
                    s.getPriority(),
                    actions,
                    templates,
                    s.getNextSteps(),
                    s.getDescription(),
                    s.getRenderType(),
                    s.getRule()
            ));
        }
        return newCache;
    }

    public Map<Long, List<com.grabathon.paymentassistant.storage.db.wrapper.Step>> buildNextStepCache(
            Map<Long, com.grabathon.paymentassistant.storage.db.wrapper.Step> newStepCache) {

        Map<Long, List<com.grabathon.paymentassistant.storage.db.wrapper.Step>> newNextStepCache = new ConcurrentHashMap<>();
        for (Map.Entry<Long, com.grabathon.paymentassistant.storage.db.wrapper.Step> entry : newStepCache.entrySet()) {
            List<com.grabathon.paymentassistant.storage.db.wrapper.Step> steps =
                    Arrays.stream(entry.getValue().getNextSteps().split(",")).filter(id -> !StringUtils.isEmpty(id)).map(id -> newStepCache.get(Long.parseLong(id))).collect(Collectors.toList());

            newNextStepCache.put(entry.getKey(), steps);
        }
        return newNextStepCache;
    }


    private void updateCache (
            Map<Long, com.grabathon.paymentassistant.storage.db.wrapper.Action> newActionCache,
            Map<Long, com.grabathon.paymentassistant.storage.db.wrapper.Template> newTemplateCache,
            Map<Long, com.grabathon.paymentassistant.storage.db.wrapper.Step> newStepCache,
            Map<Long, List<com.grabathon.paymentassistant.storage.db.wrapper.Step>> newNextStepCache) {

        lock.lock();
            actionCache.clear(); actionCache.putAll(newActionCache);
            templateCache.clear(); templateCache.putAll(newTemplateCache);
            stepCache.clear(); stepCache.putAll(newStepCache);
            nextStepCache.clear(); nextStepCache.putAll(newNextStepCache);
        lock.unlock();
    }

}
