package com.grabathon.paymentassistant.orchestration.orchestrator;

import com.grabathon.paymentassistant.orchestration.action.Action;
import com.grabathon.paymentassistant.orchestration.rules.Rules;
import com.grabathon.paymentassistant.orchestration.template.Template;
import com.grabathon.paymentassistant.storage.cache.memory.Cache;
import com.grabathon.paymentassistant.storage.db.wrapper.Step;
import com.grabathon.paymentassistant.web.api.request.NextStepsRequest;
import com.grabathon.paymentassistant.web.api.response.vo.StepResponseVO;
import com.grabathon.paymentassistant.web.api.response.vo.TemplateResponseVO;
import org.springframework.beans.factory.annotation.Autowired;
import org.springframework.stereotype.Service;
import org.springframework.util.CollectionUtils;

import java.util.LinkedList;
import java.util.List;

@Service public class Orchestrator {

    @Autowired private Rules rules;
    @Autowired private Template template;
    @Autowired private Action action;


    public OrchestratorResponse execute (Long id, NextStepsRequest request) {
        action.execute(Cache.getStepCache().get(id).getActions());
        List<Step> steps = this.getSteps(id);
        List<StepResponseVO> stepResponse = this.transformSteps(steps);
        List<com.grabathon.paymentassistant.storage.db.wrapper.Template> templates = template.get(Cache.getStepCache().get(id));
        List<TemplateResponseVO> templateResponse = this.transformTemplates(templates);
        return new OrchestratorResponse (stepResponse, templateResponse);
    }

    private List<Step> getSteps (Long id) {
        List<Step> allSteps = Cache.getNextStepCache().get(id);
        List<Step> validSteps = new LinkedList<>();
        for (Step s : allSteps) {
            if (rules.check(s)) validSteps.add(s);
        }
        return validSteps;
    }

    private List<StepResponseVO> transformSteps (List<Step> steps) {
        List<StepResponseVO> stepResponse = new LinkedList<>();
        if (CollectionUtils.isEmpty(steps)) return stepResponse;
        for (Step step : steps) {
            stepResponse.add(new StepResponseVO(step));
        }
        return stepResponse;
    }

    private List<TemplateResponseVO> transformTemplates (List<com.grabathon.paymentassistant.storage.db.wrapper.Template> templates) {
        List<TemplateResponseVO> templateResponse = new LinkedList<>();
        if (CollectionUtils.isEmpty(templates)) return templateResponse;
            for (com.grabathon.paymentassistant.storage.db.wrapper.Template template : templates) {
            templateResponse.add(new TemplateResponseVO(template));
        }
        return templateResponse;
    }

}
