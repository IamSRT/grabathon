package com.grabathon.paymentassistant.web.api.controller.utils;


import static com.grabathon.paymentassistant.utils.Utils.isEmpty;

public class Utils {


    public static Boolean notValidHeaders (
            String client,
            String user) {

        return isEmpty(client) || isEmpty(user);
    }


}
