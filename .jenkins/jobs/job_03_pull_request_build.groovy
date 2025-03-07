multibranchPipelineJob('eotc/quickfix/pull-request-build') {
  description('quickfix pull request builder.\nOwner: Yuval King, #eotc-jenkins for assistance\n')
  branchSources {
    branchSource {
      source {
        github {
          id('7519e18d-5411-4e1d-bade-895feecf366f') // IMPORTANT: use a constant and unique identifier
          apiUri('https://api.github.com')
          configuredByUrl(false)
          credentialsId('github-molly-brown-api-token')
          repoOwner('omni3x')
          repository('quickfix')
          repositoryUrl('https://github.com/omni3x/quickfix.git')
          traits {
            gitHubPullRequestDiscovery {
              strategyId(1)
            }
          }
        }
      }
      strategy {
        allBranchesSame {
          props {
            buildRetention {
              buildDiscarder {
                logRotator {
                  artifactDaysToKeepStr('7')
                  artifactNumToKeepStr('200')
                  daysToKeepStr('7')
                  numToKeepStr('200')
                }
              }
            }
          }
        }
      }
    }
  }
  factory {
    workflowBranchProjectFactory {
      scriptPath('.jenkins/jenkinsfiles/build.groovy')
    }
  }
  orphanedItemStrategy {
    defaultOrphanedItemStrategy {
      abortBuilds(false)
      daysToKeepStr('-1')
      numToKeepStr('-1')
      pruneDeadBranches(true)
    }
    discardOldItems {
      daysToKeep(-1)
      numToKeep(-1)
    }
  }
  configure {
    it << 'disabled'('false') {}
  }
}
