Feature: Create a backup
  To save backup
  As a scheduled job
  I need to save in local directory a source's backup file

  Scenario: Create a consul backup file
    Given a source "consul" equals to consul and endpoint "url" and token "token" 
    When consul backuper is executed
    Then a "backup.tgz" must be save in a local directory

  Scenario: Create a postgres backup file
    Given a source "postgres" and  "host" and "port" and "db_name" and "db_user" and "db_pass" 
    When postgres option is executed
    Then a "backup.dump" must be save in a local directory