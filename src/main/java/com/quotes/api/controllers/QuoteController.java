package com.quotes.api.controllers;

import com.quotes.api.entities.Quote;
import com.quotes.api.services.QuoteService;
import org.springframework.beans.factory.annotation.Autowired;
import org.springframework.data.domain.Sort;
import org.springframework.http.HttpStatus;
import org.springframework.http.ResponseEntity;
import org.springframework.web.bind.annotation.*;

import javax.websocket.server.PathParam;
import java.util.List;
import java.util.Optional;
import java.util.OptionalInt;

@RestController
@CrossOrigin(origins = {"http://localhost:4200", "http://quotes-web.s3-website-us-east-1.amazonaws.com"})
@RequestMapping("/api")
public class QuoteController {
    @Autowired QuoteService quoteService;

    @GetMapping("/quotes")
    public ResponseEntity<List<Quote>> GetAll(
            @RequestParam(value = "page", required = false) Optional<Integer> page,
            @RequestParam(value = "size", required = false) Optional<Integer> size,
            @RequestParam(value = "sortOrder", required = false) Optional<String> sortOrder) {
        return new ResponseEntity<>(quoteService.getAll(page.orElse(0), size.orElse(5), sortOrder.orElse("DESC")), HttpStatus.OK);
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
