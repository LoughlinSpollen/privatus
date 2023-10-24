#include <string>

#include "domain/model/FederationConfig.hpp"

#ifndef REGISTER_USECASE_H
#define REGISTER_USECASE_H

class FederationClient;
class FederationConfig;

namespace usecase {

class RegisterUsecase {
public:
    /// Constructor
    /// @param mlModel the MLModel ID that identifies the federation
    /// @param federationClient the network gRPC client
    RegisterUsecase(const std::string &mlModelId, std::shared_ptr<FederationClient> federationClient);

    /// Destructor
    ~RegisterUsecase();

    /// Get the training configuration
    /// @return federation configuration
    std::shared_ptr<FederationConfig> GetTrainingConfig() const;

private:
    const std::string m_MlModelId;
    const std::shared_ptr<FederationClient> m_FederationClient;
    std::shared_ptr<FederationConfig> m_FederationConfig;
};

}

#endif
