package com.grabathon.paymentassistant.web.api.response.vo;

import com.grabathon.paymentassistant.web.api.response.vo.base.StepResponseVO;
import lombok.Getter;
import lombok.NoArgsConstructor;
import lombok.Setter;

@Getter @Setter @NoArgsConstructor
public class RequesterStepResponseVO extends StepResponseVO {

    private String renderType;

    public RequesterStepResponseVO (StepResponseVO stepResponse) {
        this.setId(stepResponse.getId());
        this.setTitle(stepResponse.getTitle());
        this.setDescription(stepResponse.getDescription());
        this.setPriority(stepResponse.getPriority());
        this.setRenderType(stepResponse.getRequesterRenderType());
        this.setRequesteeRenderType(null);
        this.setRequesterRenderType(null);
    }

}
