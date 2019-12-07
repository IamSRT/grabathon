package com.grabathon.paymentassistant.web.api.controller;

import com.grabathon.paymentassistant.web.api.request.NextStepsRequest;
import com.grabathon.paymentassistant.web.api.request.PreviousStepsRequest;
import com.grabathon.paymentassistant.web.api.response.NextStepsResponse;
import com.grabathon.paymentassistant.web.api.response.PreviousStepsResponse;
import com.grabathon.paymentassistant.web.api.response.base.Response;
import com.grabathon.paymentassistant.web.api.routes.APIRoutes;
import com.grabathon.paymentassistant.web.api.service.PaymentAssistanceService;
import org.springframework.beans.factory.annotation.Autowired;
import org.springframework.web.bind.annotation.*;

import static com.grabathon.paymentassistant.utils.Constants.EMPTY;
import static com.grabathon.paymentassistant.web.api.controller.utils.Constants.HEADERS.CLIENT_ID;
import static com.grabathon.paymentassistant.web.api.controller.utils.Constants.HEADERS.USER_ID;
import static com.grabathon.paymentassistant.web.api.controller.utils.Constants.STEP_ID;
import static com.grabathon.paymentassistant.web.api.response.base.Code.SUCCESS;

@RestController @RequestMapping( APIRoutes.V1 ) public class StepsController {


    @Autowired private PaymentAssistanceService paymentAssistanceService;


    @PostMapping ( APIRoutes.NEXT_STEPS ) public Response<NextStepsResponse> getNextSteps (
            @RequestHeader(name = CLIENT_ID, defaultValue = EMPTY) String client,
            @RequestHeader(name = USER_ID, defaultValue = EMPTY) String user,
            @PathVariable(name = STEP_ID) Long id,
            @RequestBody NextStepsRequest request) {

        NextStepsResponse response = paymentAssistanceService.serve(id, request);
        return new Response <> (
                SUCCESS.getCode(),
                String.format(SUCCESS.getMessage(), id),
                response
        );
    }


    @PostMapping ( APIRoutes.PREVIOUS_STEPS ) public Response<PreviousStepsResponse> getPreviousSteps (
            @RequestHeader(name = CLIENT_ID, defaultValue = EMPTY) String client,
            @RequestHeader(name = USER_ID, defaultValue = EMPTY) String user,
            @PathVariable(name = STEP_ID) String id,
            @RequestBody PreviousStepsRequest request) {


        return null;
    }


}
