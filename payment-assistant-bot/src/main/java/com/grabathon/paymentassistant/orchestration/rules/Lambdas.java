package com.grabathon.paymentassistant.orchestration.rules;

import com.google.gson.Gson;
import com.grabathon.paymentassistant.http.HTTP;
import com.grabathon.paymentassistant.web.api.request.base.StepsRequest;
import com.grabathon.paymentassistant.web.api.response.wallet.Wallet;
import org.json.JSONObject;

public class Lambdas {
     private static final String walletServiceHost  = "http://localhost:4000/api/";
     private static final String Success  = "Success";

     public Float GetBalance(StepsRequest stepsRequest){
         String url = walletServiceHost + "user/" +stepsRequest.UserId;
         JSONObject response = HTTP.get(url);
         if (response == null){return (float) -1;}
         return  response.getJSONObject("wallet").getFloat("balance");
     }

    public String GetUser(StepsRequest stepsRequest){
        String url = walletServiceHost + "user/" +stepsRequest.UserId;
        JSONObject response = HTTP.get(url);
        if (response == null){return "{}";}
        return  response.toString();
    }

    public boolean Vouch(StepsRequest stepsRequest){
        String url = walletServiceHost + "user/vouch";
        String data = new Gson().toJson(stepsRequest);
        JSONObject response = HTTP.post(url,data);
        if (response == null){return false;}
        return response.getString("message").equals(Success);
    }

    public boolean EnableAutoPay(StepsRequest stepsRequest){
        String url = walletServiceHost + "user/enable-auto_pay";
        String data = new Gson().toJson(stepsRequest);
        JSONObject response = HTTP.post(url,data);
        if (response == null){return false;}
        return response.getString("message").equals(Success);
    }

    public boolean IsAutoPayEnabled(StepsRequest stepsRequest){
        String url = walletServiceHost + "user/is-auto-pay-enabled";
        String data = new Gson().toJson(stepsRequest.Vouch);
        JSONObject response = HTTP.post(url,data);
        if (response == null){return false;}
        return response.getString("message").equals(Success);
    }

    public boolean IsVouchValid(StepsRequest stepsRequest){
        String url = walletServiceHost + "user/is-vouch-valid";
        String data = new Gson().toJson(stepsRequest.Vouch);
        JSONObject response = HTTP.post(url,data);
        if (response == null){return false;}
        return response.getString("message").equals(Success);
    }

    public boolean CreatePayment(StepsRequest stepsRequest){
        String url = walletServiceHost + "payment/create";
        String data = new Gson().toJson(stepsRequest.payment);
        JSONObject response = HTTP.post(url,data);
        if (response == null){return false;}
        return response.getString("message").equals(Success);
    }

    public String GetPayment(StepsRequest stepsRequest){
        String url = walletServiceHost + "payment/"+stepsRequest.PaymentId;
        JSONObject response = HTTP.get(url);
        if (response == null){return "{}";}
        return response.toString();
    }
}
