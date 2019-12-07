package com.grabathon.paymentassistant.storage.db.entity;

import lombok.Getter;

import javax.persistence.*;

@Getter
@Entity @Table(name = "action") public class Action {

    @Id @GeneratedValue @Column (name = "id")
    private Long id;

    @Column (name = "name")
    private String name;

}
