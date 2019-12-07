package com.grabathon.paymentassistant.web.api.response;

import com.fasterxml.jackson.annotation.JsonProperty;
import com.grabathon.paymentassistant.orchestration.orchestrator.OrchestratorResponse;
import com.grabathon.paymentassistant.web.api.response.vo.RequesteeStepResponseVO;
import com.grabathon.paymentassistant.web.api.response.vo.base.StepResponseVO;
import com.grabathon.paymentassistant.web.api.response.vo.base.TemplateResponseVO;
import lombok.AllArgsConstructor;
import lombok.Getter;
import lombok.NoArgsConstructor;
import lombok.Setter;

import java.util.LinkedList;
import java.util.List;

@Getter @Setter @NoArgsConstructor @AllArgsConstructor
public class Requestee {

    @JsonProperty("options") private List<RequesteeStepResponseVO> steps;
    @JsonProperty ("messages") private List<TemplateResponseVO> templates;


    public Requestee (OrchestratorResponse orchestratorResponse) {
        this.steps = transform(orchestratorResponse.getSteps());
        this.templates = orchestratorResponse.getTemplates();
    }


    private List<RequesteeStepResponseVO> transform (List<StepResponseVO> steps) {
        List<RequesteeStepResponseVO> requesteeStepResponse = new LinkedList<>();
        for (StepResponseVO step : steps) {
            requesteeStepResponse.add(new RequesteeStepResponseVO(step));
        }
        return requesteeStepResponse;
    }

}
