#!/bin/bash

# usage: ./deployme.sh staging sddffafaafaf
# license: public domain

application_name="terena.office.service"
environment_name="Terenaofficeservice-env"

EB_BUCKET=terena.office

VERSION="v11"

aws configure set default.region ap-southeast-1

# Create a new application version with the zipped up Dockerrun file
echo "Creating new application version"
aws elasticbeanstalk create-application-version --application-name $application_name \
    --version-label $VERSION --source-bundle S3Bucket=$EB_BUCKET,S3Key=$ZIP

# Update the environment to use the new application version
echo "Updating environment"
aws elasticbeanstalk update-environment --environment-name $environment_name \
      --version-label $VERSION