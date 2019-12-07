package com.grabathon.paymentassistant.orchestration.action;

import com.grabathon.paymentassistant.orchestration.rules.Lambdas;
import com.grabathon.paymentassistant.web.api.request.base.StepsRequest;
import org.springframework.stereotype.Service;

import java.util.List;

@Service public class Action {

    public void execute (List<com.grabathon.paymentassistant.storage.db.wrapper.Action> actions, StepsRequest request) {
        for (com.grabathon.paymentassistant.storage.db.wrapper.Action action : actions) {
            switch (action.getName()) {
                case "SEND_MONEY":
                    Lambdas.CreatePayment(request);
                    break;

                case "GET_BALANCE":
                    Lambdas.GetBalance(request);
                    break;

                case "ACCEPT_VOUCH":
                    Lambdas.acceptVouch(request);
                    break;

                case "REJECT_VOUCH":
                    Lambdas.declineVouch(request);
                    break;

                case "REQUEST_VOUCH":
                    Lambdas.requestVouch(request);
                    break;

            }
        }

    }

}
