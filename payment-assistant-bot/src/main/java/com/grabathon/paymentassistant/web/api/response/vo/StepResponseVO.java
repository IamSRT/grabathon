package com.grabathon.paymentassistant.web.api.response.vo;

import com.grabathon.paymentassistant.storage.db.wrapper.Step;
import lombok.AllArgsConstructor;
import lombok.Getter;
import lombok.NoArgsConstructor;

@Getter @NoArgsConstructor @AllArgsConstructor
public class StepResponseVO {

    private Long id;
    private String title;
    private String description;
    private Integer priority;
    private String renderType;


    public StepResponseVO (Step step) {
        this.id = step.getId();
        this.title = step.getTitle();
        this.description = step.getDescription();
        this.priority = step.getPriority();
        this.renderType = step.getRenderType();
    }

}
