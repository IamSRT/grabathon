package com.grabathon.paymentassistant.web.api.controller;

import com.grabathon.paymentassistant.web.api.response.NextStepsResponse;
import com.grabathon.paymentassistant.web.api.response.base.Response;
import com.grabathon.paymentassistant.web.api.routes.APIRoutes;
import com.grabathon.paymentassistant.web.api.service.PaymentAssistanceService;
import org.apache.logging.log4j.LogManager;
import org.apache.logging.log4j.Logger;
import org.json.JSONObject;
import org.springframework.beans.factory.annotation.Autowired;
import org.springframework.web.bind.annotation.*;

import static com.grabathon.paymentassistant.web.api.controller.utils.Constants.STEP_ID;
import static com.grabathon.paymentassistant.web.api.response.base.Code.SUCCESS;

@RestController @RequestMapping( APIRoutes.V1 ) public class StepsController {

    private static final Logger logger = LogManager.getLogger(StepsController.class);

    @Autowired private PaymentAssistanceService paymentAssistanceService;


    @PostMapping ( APIRoutes.NEXT_STEPS ) public Response<NextStepsResponse> getNextSteps (
            @PathVariable(name = STEP_ID) String i,
            @RequestBody String in) {

        JSONObject request = new JSONObject(in);
        Long id = Long.parseLong(i);
        logger.info("request url:: bot.com/v1/" + i + "/next/steps");
        logger.info("request body:: " + request.toString());
        NextStepsResponse response = paymentAssistanceService.serve(id, request);
        return new Response <> (
                SUCCESS.getCode(),
                String.format(SUCCESS.getMessage(), id),
                response
        );
    }


}
