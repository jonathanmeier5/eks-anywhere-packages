#!/usr/bin/env bash

# Copyright Amazon.com Inc. or its affiliates. All Rights Reserved.
#
# Licensed under the Apache License, Version 2.0 (the "License");
# you may not use this file except in compliance with the License.
# You may obtain a copy of the License at
#
#      http://www.apache.org/licenses/LICENSE-2.0
#
# Unless required by applicable law or agreed to in writing, software
# distributed under the License is distributed on an "AS IS" BASIS,
# WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
# See the License for the specific language governing permissions and
# limitations under the License.

# Common functions used by various build scripts. To use these
# functions, source this file, e.g.:
#    . <path-to-this-file>/common.sh

# awsAuth echoes an AWS ECR password
function awsAuth () {
    local repo=${1?:no repo specified}
    local awsCmd="ecr"
    local region="--region=us-west-2"

    if [[ $repo =~ "public.ecr.aws" ]]; then
        awsCmd="ecr-public"
        region="--region=us-east-1"
    fi

    aws "$awsCmd" "$region" "--profile=${PROFILE:-}" get-login-password
}
