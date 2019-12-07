package com.grabathon.paymentassistant.web.api.response.vo;

import com.grabathon.paymentassistant.web.api.response.vo.base.StepResponseVO;
import lombok.Getter;
import lombok.NoArgsConstructor;
import lombok.Setter;

@Getter @Setter @NoArgsConstructor
public class RequesteeStepResponseVO extends StepResponseVO {

    private String renderType;

    public RequesteeStepResponseVO (StepResponseVO stepResponse) {
        this.setId(stepResponse.getId());
        this.setTitle(stepResponse.getTitle());
        this.setDescription(stepResponse.getDescription());
        this.setPriority(stepResponse.getPriority());
        this.setRenderType(stepResponse.getRequesteeRenderType());
        this.setRequesteeRenderType(null);
        this.setRequesterRenderType(null);
    }

}
