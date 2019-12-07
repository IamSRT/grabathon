package com.grabathon.paymentassistant.web.api.response.vo.base;

import com.fasterxml.jackson.annotation.JsonInclude;
import com.grabathon.paymentassistant.storage.db.wrapper.Step;
import lombok.AllArgsConstructor;
import lombok.Getter;
import lombok.NoArgsConstructor;
import lombok.Setter;

@Getter @Setter @NoArgsConstructor @AllArgsConstructor
@JsonInclude(JsonInclude.Include.NON_NULL)
public class StepResponseVO {

    private Long id;
    private String title;
    private String description;
    private Integer priority;
    private String requesterRenderType;
    private String requesteeRenderType;


    public StepResponseVO (Step step) {
        this.id = step.getId();
        this.title = step.getTitle();
        this.description = step.getDescription();
        this.priority = step.getPriority();
        this.requesterRenderType = step.getRequesterRenderType();
        this.requesteeRenderType = step.getRequesteeRenderType();
    }

}
