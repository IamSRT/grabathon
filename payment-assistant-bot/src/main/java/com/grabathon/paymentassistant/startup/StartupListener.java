package com.grabathon.paymentassistant.startup;

import com.grabathon.paymentassistant.storage.cache.memory.Cache;
import org.apache.logging.log4j.LogManager;
import org.apache.logging.log4j.Logger;
import org.springframework.beans.factory.annotation.Autowired;
import org.springframework.context.event.ContextRefreshedEvent;
import org.springframework.context.event.EventListener;
import org.springframework.stereotype.Component;

@Component public class StartupListener {

    private static final Logger logger = LogManager.getLogger(StartupListener.class);

    @Autowired private Cache cache;

    @EventListener public void handleContextRefresh (
            final ContextRefreshedEvent event) {

        try {
            cache.load();
        } catch (Exception ex) {
            logger.error(ex.getMessage(), ex);
        }
    }

}
