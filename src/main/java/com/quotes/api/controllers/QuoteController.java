package com.quotes.api.controllers;

import com.quotes.api.entities.Quote;
import com.quotes.api.services.QuoteService;

import java.util.List;
import java.util.Optional;
import java.util.concurrent.ExecutionException;

import org.springframework.beans.factory.annotation.Autowired;
import org.springframework.http.HttpStatus;
import org.springframework.http.ResponseEntity;

import org.springframework.web.bind.annotation.CrossOrigin;
import org.springframework.web.bind.annotation.DeleteMapping;
import org.springframework.web.bind.annotation.GetMapping;
import org.springframework.web.bind.annotation.PathVariable;
import org.springframework.web.bind.annotation.PostMapping;
import org.springframework.web.bind.annotation.PutMapping;
import org.springframework.web.bind.annotation.RequestBody;
import org.springframework.web.bind.annotation.RequestMapping;
import org.springframework.web.bind.annotation.RequestParam;
import org.springframework.web.bind.annotation.RestController;

@RestController
@CrossOrigin(origins = {"http://localhost:4200"})
@RequestMapping("/api")
public class QuoteController {
    @Autowired
    QuoteService quoteService;

    @GetMapping("/quotes")
    public ResponseEntity<List<Quote>> GetAll(Optional<Integer> page, Optional<Integer> size, String sortOrder) {
        try {
            List<Quote> quotes = quoteService.getAll(page, size, sortOrder);

            return new ResponseEntity<>(quotes, HttpStatus.OK);
        }
        catch (Exception exception) {
            return new ResponseEntity<>(null, HttpStatus.INTERNAL_SERVER_ERROR);
        }
    }

    @GetMapping("/quotes/search")
    public ResponseEntity<List<Quote>> search(
            @RequestParam(value = "q", required = false) String query)
        throws ExecutionException, InterruptedException {
        try{
            return new ResponseEntity<>(quoteService.findByTag(query), HttpStatus.OK);
        }
        catch (Exception exception) {
            return new ResponseEntity<>(null);
        }
    }

    @GetMapping("/quotes/{idQuote}")
    public ResponseEntity<Quote> findById(@PathVariable("idQuote") String idQuote) {
        return new ResponseEntity<>(quoteService.findById(idQuote), HttpStatus.OK);
    }

    @PostMapping("/quotes")
    public ResponseEntity<Quote> create(@RequestBody Quote quote) {
        return new ResponseEntity<>(quoteService.create(quote), HttpStatus.CREATED);
    }

    @DeleteMapping("/quotes/{idQuote}")
    public ResponseEntity<String> delete(@PathVariable("idQuote") String idQuote) {
        quoteService.delete(idQuote);
        return new ResponseEntity<>("", HttpStatus.OK);
    }

    @PutMapping("/quotes")
    public ResponseEntity<Quote> update(@RequestBody Quote quote) {
        return new ResponseEntity<>(quoteService.update(quote), HttpStatus.OK);
    }
}
