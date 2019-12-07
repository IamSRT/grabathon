package com.grabathon.paymentassistant.orchestration.rules;

import com.google.gson.Gson;
import com.grabathon.paymentassistant.http.HTTP;
import com.grabathon.paymentassistant.web.api.request.base.StepsRequest;
import com.grabathon.paymentassistant.web.api.request.base.Vouch;
import org.json.JSONObject;

import java.util.LinkedList;
import java.util.List;

public class Lambdas {
     private static final String walletServiceHost  = "http://192.168.43.141:4000/api/";
     private static final String Success  = "Success";

     public static Float GetBalance(StepsRequest stepsRequest){
         String url = walletServiceHost + "user/" +stepsRequest.requesterNumber;
         JSONObject response = HTTP.get(url);
         if (response == null){return (float) -1;}
         return  response.getJSONObject("data").getJSONObject("wallet").getFloat("balance");
     }

    public static String GetUser(StepsRequest stepsRequest){
        String url = walletServiceHost + "user/" +stepsRequest.UserId;
        JSONObject response = HTTP.get(url);
        if (response == null){return "{}";}
        return  response.toString();
    }

    public static boolean requestVouch (StepsRequest stepsRequest){
        String url = walletServiceHost + "user/vouch";
        List<Vouch> vouches = new LinkedList<>();
        for (String v : stepsRequest.requesteeNumber) {
            Vouch vouch = new Vouch();
            vouch.VoucheeId = stepsRequest.requesterNumber;
            vouch.VoucherId = v;
            vouch.Amount = stepsRequest.amount;
            vouches.add(vouch);
        }
        String data = new Gson().toJson(stepsRequest);
        JSONObject response = HTTP.post(url,data);
        if (response == null){return false;}
        return response.getString("message").equals(Success);
    }

    public static boolean acceptVouch(StepsRequest stepsRequest){
        String url = walletServiceHost + "user/vouch" + stepsRequest.vouch_id + "/accept";
        String data = new Gson().toJson(stepsRequest);
        JSONObject response = HTTP.post(url,data);
        if (response == null){return false;}
        return response.getString("message").equals(Success);
    }

    public static boolean declineVouch(StepsRequest stepsRequest){
        String url = walletServiceHost + "user/vouch/" + stepsRequest.vouch_id + "/reject";
        String data = new Gson().toJson(stepsRequest);
        JSONObject response = HTTP.post(url,data);
        if (response == null){return false;}
        return response.getString("message").equals(Success);
    }

    public static boolean EnableAutoPay(StepsRequest stepsRequest){
        String url = walletServiceHost + "user/enable-auto_pay";
        String data = new Gson().toJson(stepsRequest);
        JSONObject response = HTTP.post(url,data);
        if (response == null){return false;}
        return response.getString("message").equals(Success);
    }

    public static boolean IsAutoPayEnabled(StepsRequest stepsRequest){
        String url = walletServiceHost + "user/is-auto-pay-enabled";
        String data = new Gson().toJson(stepsRequest.Vouch);
        JSONObject response = HTTP.post(url,data);
        if (response == null){return false;}
        return response.getString("message").equals(Success);
    }

    public static boolean IsVouchValid(StepsRequest stepsRequest){
        String url = walletServiceHost + "user/is-vouch-valid";
        String data = new Gson().toJson(stepsRequest.Vouch);
        JSONObject response = HTTP.post(url,data);
        if (response == null){return false;}
        return response.getString("message").equals(Success);
    }

    public static boolean CreatePayment(StepsRequest stepsRequest){
        String url = walletServiceHost + "payment/create";
        String data = new Gson().toJson(stepsRequest.payment);
        JSONObject response = HTTP.post(url,data);
        if (response == null){return false;}
        return response.getString("message").equals(Success);
    }

    public static String GetPayment(StepsRequest stepsRequest){
        String url = walletServiceHost + "payment/"+stepsRequest.PaymentId;
        JSONObject response = HTTP.get(url);
        if (response == null){return "{}";}
        return response.toString();
    }
}
