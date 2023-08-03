import socket
import ssl

hostname = 'example.com'
context = ssl.create_default_context()

with socket.create_connection((hostname, 443)) as sock:
    with context.wrap_socket(sock, server_hostname=hostname) as ssock:
        print(ssock.version())

        # Send an HTTP request over the encrypted connection
        ssock.sendall(b"GET / HTTP/1.1\r\nHost: example.com\r\n\r\n")

        # Receive the HTTP response over the encrypted connection
        response = ssock.recv()

        # Close the connection
        sock.close()

        print(response)
