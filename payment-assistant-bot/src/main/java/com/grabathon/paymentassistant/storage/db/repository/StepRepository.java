package com.grabathon.paymentassistant.storage.db.repository;

import com.grabathon.paymentassistant.storage.db.entity.Step;
import org.springframework.data.jpa.repository.JpaRepository;
import org.springframework.stereotype.Repository;

@Repository
public interface StepRepository extends JpaRepository<Step, Long> { }
