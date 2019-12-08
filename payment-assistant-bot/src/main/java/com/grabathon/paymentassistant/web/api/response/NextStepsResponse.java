package com.grabathon.paymentassistant.web.api.response;

import com.grabathon.paymentassistant.orchestration.orchestrator.OrchestratorResponse;
import com.grabathon.paymentassistant.web.api.response.base.StepResponse;
import lombok.ToString;

@ToString
public class NextStepsResponse extends StepResponse {

    public NextStepsResponse (OrchestratorResponse orchestratorResponse) {
        this.setRequestee(new Requestee(orchestratorResponse));
        this.setRequester(new Requester(orchestratorResponse));
        this.setData(orchestratorResponse.getData());
    }

}
