#!/usr/bin/env python
import os
import sys

import pika

def callback(ch, method, properties, body):
    print(f" [x] Received {body.decode()}")

def main():
    print("connection to rabbitmq")
    connection = pika.BlockingConnection(
        pika.ConnectionParameters(host="localhost"),
    )
    print("connected")
    channel = connection.channel()

    channel.queue_declare(queue="hello")

    print("channel declared")

    def callback(ch, method, properties, body):
        print(f" [x] Received {body.decode()}")

    channel.basic_consume(
        queue="hello",
        on_message_callback=callback,
        auto_ack=True,
    )
    print(" [*] Waiting for messages. To exit press CTRL+C")
    channel.start_consuming()

    

    

if __name__ == "__main__":
    try:
        print("starting server")
        main()
    except KeyboardInterrupt:
        print("Interrupted")
        try:
            sys.exit(0)
        except SystemExit:
            os._exit(0)