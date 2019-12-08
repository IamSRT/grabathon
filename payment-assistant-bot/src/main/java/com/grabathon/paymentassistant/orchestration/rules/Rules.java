package com.grabathon.paymentassistant.orchestration.rules;

import com.grabathon.paymentassistant.storage.db.wrapper.Step;
import org.json.JSONObject;
import org.springframework.stereotype.Service;

@Service public class Rules {

    public boolean check (Step step, JSONObject request) {
        switch (step.getRule()) {
            case "BALANCE_CHECK":
                return Lambdas.GetBalance(request) > 0.0;

            case "AUTO_PAY_NOT_ENABLED":
                return !Lambdas.IsAutoPayEnabled(request);
        }
        return true;
    }

}
