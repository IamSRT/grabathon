package com.grabathon.paymentassistant.orchestration.template;

import com.grabathon.paymentassistant.storage.db.wrapper.Step;
import org.springframework.stereotype.Service;
import org.springframework.util.CollectionUtils;

import java.util.LinkedList;
import java.util.List;

@Service public class Template {

    public List<com.grabathon.paymentassistant.storage.db.wrapper.Template> get (Step step) {
        if (CollectionUtils.isEmpty(step.getTemplates())) return new LinkedList<>();
        return step.getTemplates();
    }

}
