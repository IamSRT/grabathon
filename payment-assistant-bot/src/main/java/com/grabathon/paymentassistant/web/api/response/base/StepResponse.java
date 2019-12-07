package com.grabathon.paymentassistant.web.api.response.base;

import com.grabathon.paymentassistant.web.api.response.Requestee;
import com.grabathon.paymentassistant.web.api.response.Requester;
import lombok.AllArgsConstructor;
import lombok.Getter;
import lombok.NoArgsConstructor;
import lombok.Setter;

@Getter @Setter @NoArgsConstructor @AllArgsConstructor
public class StepResponse {

    private Requester requester;
    private Requestee requestee;

}
