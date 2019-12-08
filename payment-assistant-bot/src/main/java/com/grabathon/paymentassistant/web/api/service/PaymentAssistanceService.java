package com.grabathon.paymentassistant.web.api.service;

import com.grabathon.paymentassistant.orchestration.orchestrator.Orchestrator;
import com.grabathon.paymentassistant.orchestration.orchestrator.OrchestratorResponse;
import com.grabathon.paymentassistant.web.api.response.NextStepsResponse;
import org.json.JSONObject;
import org.springframework.beans.factory.annotation.Autowired;
import org.springframework.stereotype.Service;

@Service public class PaymentAssistanceService {

    @Autowired private Orchestrator orchestrator;

    public NextStepsResponse serve (Long id, JSONObject request) {
        OrchestratorResponse orchestratorResponse = orchestrator.execute(id, request);
        return new NextStepsResponse (orchestratorResponse);
    }

}
