import random
import sys
import logging
import time
import xtuples_jobs as jobs
import os


if __name__ == '__main__':
    def job_factory(name):
        def job(key):
            """job execution function"""
            # Wait a random number of seconds between 1 and 10
            wait_time = random.randint(1, 10)
            time.sleep(wait_time)
            chance = random.randint(1, 10)
            ret = 0 if chance == 5 else 1

            # Log the name and key
            logging.info('exec job name: %s, key: %s', name, key)
            return ret
        return job

    if 'job_name' in os.environ:
        job_name = os.environ['job_name']
    else:
        logging.fatal("Error: job_name environment variable not provided.")
        sys.exit(1)

    jobs.run_server(job_name, job_factory)

