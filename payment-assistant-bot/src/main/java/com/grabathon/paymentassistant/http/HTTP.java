package com.grabathon.paymentassistant.http;

import org.json.JSONObject;
import org.springframework.http.HttpEntity;
import org.springframework.http.HttpHeaders;
import org.springframework.http.HttpMethod;
import org.springframework.http.ResponseEntity;
import org.springframework.web.client.RestClientException;
import org.springframework.web.client.RestTemplate;

public class HTTP {

    public static JSONObject get (String url) {
        RestTemplate restTemplate = new RestTemplate();
        ResponseEntity<String> response;
        try {
            HttpHeaders headers = new HttpHeaders();
            HttpEntity<String> entity = new HttpEntity<>(headers);
            response = restTemplate.exchange (
                    url,
                    HttpMethod.GET,
                    entity,
                    String.class
            );
        } catch (RestClientException rce) {
            return null;
        }
        return parseResponse(response);
    }

    public static JSONObject post (String url, String data) {
        RestTemplate restTemplate = new RestTemplate();
        ResponseEntity<String> response;
        try {
            HttpHeaders headers = new HttpHeaders();
            HttpEntity<String> entity = new HttpEntity<>(data, headers);
            response = restTemplate.exchange (
                    url,
                    HttpMethod.POST,
                    entity,
                    String.class
            );
        } catch (RestClientException rce) {
            return null;
        }
        return parseResponse(response);
    }



    private static JSONObject parseResponse (ResponseEntity<String> response) {
        return new JSONObject(response.getBody());
    }

}
