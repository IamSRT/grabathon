package com.grabathon.paymentassistant.storage.db.repository;

import com.grabathon.paymentassistant.storage.db.entity.Action;
import org.springframework.data.jpa.repository.JpaRepository;
import org.springframework.stereotype.Repository;


@Repository
public interface ActionRepository extends JpaRepository<Action, Long> { }
