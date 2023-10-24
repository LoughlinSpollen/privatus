#include "../../../include/domain/model/FederationConfig.hpp"

#include "../../../include/infra/network/FederationClient.hpp"
#include "../../../include/usecase/RegisterUsecase.hpp"

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
