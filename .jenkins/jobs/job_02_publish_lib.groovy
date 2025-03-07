pipelineJob('eotc/quickfix/publish') {
  description('tag quickfix release')
  parameters {
    stringParam('branch_name', 'master', 'Branch from quickfix repo')
  }
  definition {
    cpsScm {
      scm {
        git {
          remote {
            github('omni3x/quickfix', 'ssh')
            credentials('github-molly-brown-ssh-key')
            refspec('+refs/heads/${branch_name}:refs/remotes/origin/${branch_name}')
          }
          branch('origin/${branch_name}')
          extensions {
            cloneOptions {
              depth(1)
              honorRefspec(true)
              noTags(true)
              shallow(true)
              reference('/var/lib/gitcache/quickfix.git')
            }
          }
        }
      }
      scriptPath('.jenkins/jenkinsfiles/build_and_publish.groovy')
    }
  }
  logRotator {
    artifactDaysToKeep(90)
    artifactNumToKeep(150)
    daysToKeep(90)
    numToKeep(150)
  }
}
