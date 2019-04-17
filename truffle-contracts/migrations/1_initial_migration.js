const Migrations = artifacts.require("givo");

module.exports = function(deployer) {
  deployer.deploy(Migrations);
};
