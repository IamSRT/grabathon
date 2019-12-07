package com.grabathon.paymentassistant.web.api.request.context;

import java.util.Objects;

public class RequestContext {

    private static InheritableThreadLocal<RequestContext> context = new InheritableThreadLocal<>();


    private String requestId;
    public String getRequestId () { return requestId; }
    public void setRequestId (String requestId) { this.requestId = requestId; }

    private String tenant;
    public String getTenant () { return tenant; }
    public void setTenant (String tenant) { this.tenant = tenant; }

    private String client;
    public String getClient () { return client; }
    public void setClient (String client) { this.client = client; }

    private String entityId;
    public String getEntityId() { return entityId; }
    public void setEntityId(String entityId) { this.entityId = entityId; }



    public static RequestContext current () {
        RequestContext currentContext =  context.get();
        if (Objects.isNull(currentContext)) {
            currentContext = new RequestContext();
            context.set(currentContext);
        }
        return currentContext;
    }

}
