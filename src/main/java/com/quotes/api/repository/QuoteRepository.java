package com.quotes.api.repository;

import com.quotes.api.entities.Quote;
import org.springframework.data.mongodb.repository.MongoRepository;

import java.util.List;

public interface QuoteRepository extends MongoRepository<Quote, String> {
    List<Quote> findByTagsContaining(String[] tag);
    List<Quote> findByAuthorContaining(String author);
    List<Quote> findByWorkContaining(String work);
    List<Quote> findByMessageContaining(String message);
}