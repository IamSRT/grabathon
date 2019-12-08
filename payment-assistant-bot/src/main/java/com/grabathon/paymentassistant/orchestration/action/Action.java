package com.grabathon.paymentassistant.orchestration.action;

import com.grabathon.paymentassistant.orchestration.rules.Lambdas;
import org.json.JSONObject;
import org.springframework.stereotype.Service;

import java.util.List;

@Service public class Action {

    public Object execute (List<com.grabathon.paymentassistant.storage.db.wrapper.Action> actions, JSONObject request) {
        for (com.grabathon.paymentassistant.storage.db.wrapper.Action action : actions) {
            switch (action.getName()) {
                case "SEND_MONEY":
                case "ADD_MONEY":
                    return Lambdas.CreatePayment(request);

                case "GET_BALANCE":
                    return Lambdas.GetBalance(request);

                case "ACCEPT_VOUCH":
                    return Lambdas.acceptVouch(request);

                case "REJECT_VOUCH":
                    return Lambdas.declineVouch(request);

                case "REQUEST_VOUCH":
                    return Lambdas.requestVouch(request);

                case "GET_PENDING_PAYMENTS":
                    return Lambdas.GetPendingPayment(request);

            }
        }
        return null;
    }

}
