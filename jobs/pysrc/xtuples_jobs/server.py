import logging

from concurrent.futures import ThreadPoolExecutor

import grpc
from .xtuples_pb2 import  JobResponse
from .xtuples_pb2_grpc import JobServiceServicer, add_JobServiceServicer_to_server

# call the job function from its name for a specific key
class JobServer(JobServiceServicer):
    def __init__(self, job_name, job_factory):
        self.job_name = job_name
        self.job_factory = job_factory(job_name)
            
    def CallJob(self, request, context):
        logging.info('call job name: %s and key : %s ', self.name, request.key)
        exec_ok = self.job_factory(request.key)
        logging.info('execution of the job results in %s', 'success' if exec_ok  else 'failure')
        resp = JobResponse(ret=exec_ok)
        return resp

def run_server(job_name, job_factory):
    logging.basicConfig(
        level=logging.INFO,
        format='%(asctime)s - %(levelname)s - %(message)s',
    )
    server = grpc.server(ThreadPoolExecutor())
    add_JobServiceServicer_to_server(JobServer(job_name, job_factory), server)
    port = 9999
    server.add_insecure_port(f'[::1]:{port}')
    server.start()
    logging.info('server ready on port %r', port)
    server.wait_for_termination()

# if __name__ == '__main__':