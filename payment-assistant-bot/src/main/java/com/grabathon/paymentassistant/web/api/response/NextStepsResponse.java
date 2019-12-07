package com.grabathon.paymentassistant.web.api.response;

import com.grabathon.paymentassistant.orchestration.orchestrator.OrchestratorResponse;
import com.grabathon.paymentassistant.web.api.response.base.StepResponse;

public class NextStepsResponse extends StepResponse {

    public NextStepsResponse (OrchestratorResponse orchestratorResponse) {
        this.setSteps(orchestratorResponse.getSteps());
        this.setTemplates(orchestratorResponse.getTemplates());
    }

}
