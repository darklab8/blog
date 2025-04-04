# Intro

When we use ephemeral CI runners for docker building and running, we often have a problem to speed up the process with some caching. Docker building stage to install all dependencies can reach good 20 minutes and that could be taking 50% of CI time.

We could be using CI native caching solutions, but their usage is... not really applicable for trully stateless autoscaling CI runners, like for the used [Philips Labs Terraform AWS Github Runners](<https://github.com/philips-labs/terraform-aws-github-runner>) configuration.

How can we resolve it? There is a neat recipe usable in the same way pretty much for any type of CI. It involves building image layers involving dependency installation and storing them temporally


1) Calculate md5 hash from dependencies file lock onto libraries. Whatever your language is. Lets call the value as builder_base_hash
2) Pull image image {{ builder_base_hash }} if it exists in docker registry, if not then build up to stage --builder. Save result to docker registry under tag  {{ builder_base_hash }}
(Thus we implemented speed up twice for CI), as we are were able to cache half of longest CI in a way that its CI jobs can run at different runners, as we use remote for persistence
3) run full building of an image, til the code capable to run unit tests and push to docker registry under tag  build_${{ github.run_id }}
4) at unit test stage: pull the image  build_${{ github.run_id }} and run unit tests and other tests
5) if it passed them, than save the image as service_name_{{ github.run_number }} as fit for deployment 🙂, also mark it as latest and etc whatever tags u need 
