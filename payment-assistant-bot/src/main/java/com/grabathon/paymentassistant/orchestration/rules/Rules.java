package com.grabathon.paymentassistant.orchestration.rules;

import com.grabathon.paymentassistant.storage.db.wrapper.Step;
import org.springframework.stereotype.Service;

@Service public class Rules {

    public boolean check (Step step) {
        return true;
    }

}
