package com.amazonaws.ml.mms.metrics;

import java.io.IOException;
import java.net.InetSocketAddress;
import java.nio.ByteBuffer;
import java.nio.channels.DatagramChannel;
import java.nio.charset.StandardCharsets;
import org.slf4j.Logger;
import org.slf4j.LoggerFactory;

public class StatsDClient {
    private static final Logger logger = LoggerFactory.getLogger(StatsDClient.class);
    private String hostName;
    private int portNumber;

    public StatsDClient(String hostName, int portNumber) {
        this.hostName = hostName;
        this.portNumber = portNumber;
    }

    public void send(String metricMessage) {
        try {
            DatagramChannel clientSocket;
            clientSocket = DatagramChannel.open();
            clientSocket.connect(new InetSocketAddress(hostName, portNumber));
            byte[] byteArr = metricMessage.getBytes(StandardCharsets.UTF_8);
            clientSocket.write(ByteBuffer.wrap(byteArr));
            clientSocket.close();
        } catch (IOException e) {
            logger.error("", e);
        }
    }
}
