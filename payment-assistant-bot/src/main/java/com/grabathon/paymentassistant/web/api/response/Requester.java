package com.grabathon.paymentassistant.web.api.response;

import com.fasterxml.jackson.annotation.JsonProperty;
import com.grabathon.paymentassistant.orchestration.orchestrator.OrchestratorResponse;
import com.grabathon.paymentassistant.web.api.response.vo.RequesterStepResponseVO;
import com.grabathon.paymentassistant.web.api.response.vo.base.StepResponseVO;
import com.grabathon.paymentassistant.web.api.response.vo.base.TemplateResponseVO;
import lombok.AllArgsConstructor;
import lombok.Getter;
import lombok.NoArgsConstructor;
import lombok.Setter;

import java.util.LinkedList;
import java.util.List;

@Getter @Setter @NoArgsConstructor @AllArgsConstructor
public class Requester {

    @JsonProperty("options") private List<RequesterStepResponseVO> steps;
    @JsonProperty ("messages") private List<TemplateResponseVO> templates;


    public Requester (OrchestratorResponse orchestratorResponse) {
        this.steps = transformX(orchestratorResponse.getSteps());
        this.templates = orchestratorResponse.getTemplates();
    }


    private List<RequesterStepResponseVO> transformX (List<StepResponseVO> steps) {
        List<RequesterStepResponseVO> requesteeStepResponse = new LinkedList<>();
        for (StepResponseVO step : steps) {
            requesteeStepResponse.add(new RequesterStepResponseVO(step));
        }
        return requesteeStepResponse;
    }

}
