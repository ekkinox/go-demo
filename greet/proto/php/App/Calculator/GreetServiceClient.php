<?php
// GENERATED CODE -- DO NOT EDIT!

namespace App\Calculator;

/**
 */
class GreetServiceClient extends \Grpc\BaseStub {

    /**
     * @param string $hostname hostname
     * @param array $opts channel options
     * @param \Grpc\Channel $channel (optional) re-use channel object
     */
    public function __construct($hostname, $opts, $channel = null) {
        parent::__construct($hostname, $opts, $channel);
    }

    /**
     * Unary
     * @param \App\Calculator\GreetRequest $argument input argument
     * @param array $metadata metadata
     * @param array $options call options
     * @return \Grpc\UnaryCall
     */
    public function Greet(\App\Calculator\GreetRequest $argument,
      $metadata = [], $options = []) {
        return $this->_simpleRequest('/greet.GreetService/Greet',
        $argument,
        ['\App\Calculator\GreetResponse', 'decode'],
        $metadata, $options);
    }

    /**
     * Server streaming
     * @param \App\Calculator\GreetManyTimesRequest $argument input argument
     * @param array $metadata metadata
     * @param array $options call options
     * @return \Grpc\ServerStreamingCall
     */
    public function GreetManyTimes(\App\Calculator\GreetManyTimesRequest $argument,
      $metadata = [], $options = []) {
        return $this->_serverStreamRequest('/greet.GreetService/GreetManyTimes',
        $argument,
        ['\App\Calculator\GreetManyTimesResponse', 'decode'],
        $metadata, $options);
    }

    /**
     * Client streaming
     * @param array $metadata metadata
     * @param array $options call options
     * @return \Grpc\ClientStreamingCall
     */
    public function LongGreet($metadata = [], $options = []) {
        return $this->_clientStreamRequest('/greet.GreetService/LongGreet',
        ['\App\Calculator\LongGreetResponse','decode'],
        $metadata, $options);
    }

    /**
     * BiDi streaming
     * @param array $metadata metadata
     * @param array $options call options
     * @return \Grpc\BidiStreamingCall
     */
    public function GreetAll($metadata = [], $options = []) {
        return $this->_bidiRequest('/greet.GreetService/GreetAll',
        ['\App\Calculator\GreetAllResponse','decode'],
        $metadata, $options);
    }

}
