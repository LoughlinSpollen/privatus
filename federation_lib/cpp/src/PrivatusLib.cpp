
namespace privatus {

PrivatusLibrary::PrivatusLibrary(std::string mlModelId, const FederationStateChangeObserver &observer) :
    m_FederationClient(std::make_shared<FederationClient>(observer)),
    m_RegisterUsecase(std::make_shared<usecase::RegisterUsecase(mlModelId, m_FederationClient))
{

}

PrivatusLibrary::~PrivatusLibrary() {

}

domain::FederationConfig PrivatusLibrary::GetTrainingConfig() const {
    return m_RegisterUsecase->GetFederationConfig();
}

void PrivatusLibrary::JoinFederation() {
    m_RegisterUsecase->Register();
}

}
