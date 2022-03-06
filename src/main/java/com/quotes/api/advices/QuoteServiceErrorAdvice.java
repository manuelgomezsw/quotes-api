package com.quotes.api.advices;

import com.quotes.api.exceptions.QuoteBadRequestException;
import com.quotes.api.exceptions.QuoteNotFoundException;
import org.apache.logging.log4j.LogManager;
import org.apache.logging.log4j.Logger;
import org.springframework.http.HttpStatus;
import org.springframework.web.bind.annotation.ControllerAdvice;
import org.springframework.web.bind.annotation.ExceptionHandler;
import org.springframework.web.bind.annotation.ResponseStatus;

import java.util.concurrent.ExecutionException;

@ControllerAdvice
public class QuoteServiceErrorAdvice {
    private static final Logger logger = LogManager.getLogger(QuoteServiceErrorAdvice.class);

    @ResponseStatus(HttpStatus.NOT_FOUND)
    @ExceptionHandler({QuoteNotFoundException.class})
    public void handle(QuoteNotFoundException exception) {}

    @ResponseStatus(HttpStatus.BAD_REQUEST)
    @ExceptionHandler({QuoteBadRequestException.class})
    public void handle(QuoteBadRequestException exception) {}

    @ResponseStatus(HttpStatus.INTERNAL_SERVER_ERROR)
    @ExceptionHandler({NullPointerException.class, ExecutionException.class, InterruptedException.class})
    public void handle() {}
}
