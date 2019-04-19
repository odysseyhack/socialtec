const givo_contracts = artifacts.require("givo");

module.exports = function(deployer) {
  deployer.deploy(givo_contracts);
};
