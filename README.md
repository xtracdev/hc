## HC

Health check for docker containers. Allow access to docker container health check information without remote access to the docker daemon.

To run, map /var/run as a volume to allow access to docker.sock, and map
the container port 5000 to whatever port on the host you'd like to
access the health check at.

Note the container will need to implement the docker HEALTCHECK to
give a true picture of health.

<pre>
docker run --volume=/var/run:/var/run:rw -p 15000:5000 xtracdev/hc
</pre>

Access the endpoint using GET

<pre>
curl localhost:15000/health
/zen_visvesvaraya xtracdev/hc: Up 3 seconds
/silly_bohr sath89/oracle-12c: Up 44 minutes
/infallible_raman echo: Up 45 minutes (unhealthy)
</pre>

<pre>
curl localhost:15000/health/infallible_raman
</prev>


