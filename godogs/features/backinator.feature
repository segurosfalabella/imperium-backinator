Feature: Create consul backup
  To save consul snapshot
  As a scheduled job
  I need to be able to get snapshot from API

  Scenario: Get snapshot from consul
    Given an endpoint
    When I call get snapshot
    Then I get snapshot backup file named "backup.tgz"

  Scenario: Upload snapshot to cloud storage
    Given an snapshot
    When I call cloud storage endpoint
    Then I put backup file in cloud storage
