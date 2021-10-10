package com.quotes.api.controllers;

import com.quotes.api.entities.Quote;
import com.quotes.api.services.QuoteService;
import org.springframework.beans.factory.annotation.Autowired;
import org.springframework.http.HttpStatus;
import org.springframework.http.ResponseEntity;
import org.springframework.web.bind.annotation.*;

import java.util.List;

@RestController
@CrossOrigin(origins = "http://localhost:4200")
@RequestMapping("/api")
public class QuoteController {
    @Autowired QuoteService quoteService;

    @GetMapping("/quotes")
    public ResponseEntity<List<Quote>> GetAll() {
        return new ResponseEntity<>(quoteService.getAll(), HttpStatus.OK);
    }

    @GetMapping("/quotes/{idQuote}")
    public ResponseEntity<Quote> FindById(@PathVariable("idQuote") String idQuote) {
        return new ResponseEntity<>(quoteService.findById(idQuote), HttpStatus.OK);
    }

    @PostMapping("/quotes")
    public ResponseEntity<Quote> Create(@RequestBody Quote quote) {
        return new ResponseEntity<>(quoteService.create(quote), HttpStatus.CREATED);
    }

    @DeleteMapping("/quotes/{idQuote}")
    public ResponseEntity<String> Delete(@PathVariable("idQuote") String idQuote) {
        quoteService.delete(idQuote);
        return new ResponseEntity<>("", HttpStatus.OK);
    }

    @PutMapping("/quotes")
    public ResponseEntity<Quote> Update(@RequestBody Quote quote) {
        return new ResponseEntity<>(quoteService.update(quote), HttpStatus.OK);
    }
}
