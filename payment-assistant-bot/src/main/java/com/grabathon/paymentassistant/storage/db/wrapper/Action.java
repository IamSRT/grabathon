package com.grabathon.paymentassistant.storage.db.wrapper;

import lombok.AllArgsConstructor;
import lombok.Getter;
import lombok.NoArgsConstructor;

@Getter @NoArgsConstructor @AllArgsConstructor
public class Action {

    private Long id;

    private String name;

    public Action (com.grabathon.paymentassistant.storage.db.entity.Action action) {
        this.id = action.getId();
        this.name = action.getName();
    }

}
