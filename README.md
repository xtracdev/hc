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

Access the endpoint using GET - /health returns the health status of all
containers, or get the health of a specific container via container name 
(/health/<container-name>).

<pre>
curl localhost:15000/health
/modest_cori echo: Up 3 minutes (healthy)
/silly_bohr sath89/oracle-12c: Up 55 minutes
</pre>

Note the container will need to implement the docker HEALTCHECK to
give a true picture of health. Without this we can only glean how long
something has beed up, not if it is healthy or not.

Example reflecting positive health status.

<pre>
curl localhost:15000/health/modest_cori
/modest_cori echo: Up 4 minutes (healthy)
</pre>

Example reflecting negative health status

<pre>
curl localhost:15000/health/modest_cori
/modest_cori echo: Up 6 minutes (unhealthy)
</pre>
