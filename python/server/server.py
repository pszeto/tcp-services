import socket
import logging
import os
import threading

def handle_client(client_socket):
    podName = os.environ.get('POD_NAME')
    podNamespace = os.environ.get('POD_NAMESPACE')
    while True:
        data = client_socket.recv(1024).decode()
        if not data:
            logging.info("No data from connected user -- possibly disconnected ")
            break
        logging.info("Message from connected user: " + str(data))
        data = ("(" + podName + "." + podNamespace + ") Echoing back: " + str(data))
        client_socket.send(data.encode())  # send data to the client
        logging.info("Sent: " + str(data))
    client_socket.close()

def run_server():
    logging.basicConfig(format='%(asctime)s %(levelname)-8s %(message)s',level=logging.INFO,datefmt='%Y-%m-%d %H:%M:%S')

    # get the hostname
    host = "0.0.0.0"
    if "PORT" in os.environ:
        port = int(os.environ.get('PORT'))
    else:
        port = 5000

    server_socket = socket.socket(socket.AF_INET, socket.SOCK_STREAM)
    server_socket.bind((host, port))
    server_socket.listen(10)

    logging.info("Successfully bound :" + str(host) + ":" + str(port))

    while True:
        client_socket, address = server_socket.accept()
        logging.info("Connection from: " + str(address))
        thread = threading.Thread(target=handle_client, args=(client_socket,))
        thread.start()

if __name__ == "__main__":
    run_server()
