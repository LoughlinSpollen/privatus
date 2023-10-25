#include "RegisterUsecase.hpp"

#include <string>
#include <memory>

#include "FederationConfig.hpp"
#include "FederationClient.hpp"

namespace usecase {

RegisterUsecase::RegisterUsecase(const std::string &modelId, std::shared_ptr<FederationClient> federationClient) : 
    m_MlModelId(modelId),
    m_FederationClient(federationClient) {

}

RegisterUsecase::~RegisterUsecase() {

}

std::shared_ptr<FederationConfig> RegisterUsecase::GetTrainingConfig() const {
    return m_FederationConfig;
}

}
