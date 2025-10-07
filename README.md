# Kiln Monitoring Microservice

## 🎯 Vue d'ensemble

Ce microservice Go surveille en temps réel les performances des validateurs Kiln en collectant périodiquement les métriques de récompenses, uptime et statut. Il applique des règles configurables et peut envoyer des alertes via Discord ou stdout.

## 🏗️ Architecture

```
┌─────────────────┐    ┌─────────────────┐    ┌─────────────────┐
│   Config YAML   │───▶│   Collector     │───▶│  Kiln API       │
│   (Règles)      │    │   (Scheduler)   │    │  (Rewards)      │
└─────────────────┘    └─────────────────┘    └─────────────────┘
                                │
                                ▼
                       ┌─────────────────┐
                       │   Alert System  │
                       │   (Discord/stdout) │
                       └─────────────────┘
```

## 🚀 Fonctionnalités

- **Monitoring périodique** : Collecte automatique des métriques selon un intervalle configurable
- **Mode démo** : Fonctionne sans token API avec des données simulées
- **Configuration flexible** : YAML + variables d'environnement
- **Système d'alertes** : Support Discord webhook et stdout
- **Health check** : Endpoint `/health` pour monitoring
- **Graceful shutdown** : Arrêt propre avec signal handling

## 📋 Prérequis

- Go 1.19+
- Docker (optionnel)

## 🛠️ Installation et Configuration

### 1. Cloner le projet
```bash
git clone <repository-url>
cd Kiln-monitoring-microservice
```

### 2. Configuration
Copiez et modifiez le fichier de configuration :
```bash
cp configs/config.example.yaml configs/config.yaml
```

### 3. Variables d'environnement (optionnel)
```bash
export KILN_API_TOKEN="your_actual_token"
export DISCORD_WEBHOOK_URL="https://discord.com/api/webhooks/..."
export ALERT_KIND="discord"
```

## 🏃‍♂️ Utilisation

### Mode Développement
```bash
# Installer les dépendances
go mod tidy

# Lancer le service
go run main.go
```

### Tests
```bash
# Tests unitaires
go test ./internal/kiln

# Test du client seul
go run internal/test-client.go
```

### Docker (à venir)
```bash
docker build -t kiln-monitor .
docker run -p 8080:8080 kiln-monitor
```

## 📊 Configuration

### Structure du fichier `config.yaml`
```yaml
kiln:
  baseURL: https://api.kiln.fi
  token: ""                    # Vide = mode démo

monitoring:
  interval: 60s               # Fréquence de collecte
  account_ids: ["acc_demo"]   # Comptes à surveiller

alerts:
  kind: stdout               # stdout | discord
  discord_webhook: ""

rules:
  min_uptime: 0.995          # 99.5% minimum
  max_inactive_periods: 2
  reward_drop_threshold_pct: 40
```

## 🔍 API Endpoints

### Health Check
```bash
curl http://localhost:8080/health
```
**Réponse :**
```json
{"status":"ok","service":"kiln-monitoring"}
```

## 📈 Métriques Collectées

Le service collecte les données suivantes pour chaque compte :

- **AccountID** : Identifiant du compte
- **TotalReward** : Récompenses en ETH sur la période
- **Uptime** : Pourcentage de disponibilité (0.0 - 1.0)
- **Status** : Statut du validateur (active/inactive)
- **WindowStart/End** : Période de collecte

## 🧪 Tests et Validation

### Test du client Kiln
```bash
go run internal/test-client.go
```
**Sortie attendue :**
```
=== Test du client Kiln ===
Account ID: acc_demo
Période: 14:30 → 15:30
Récompense: 0.023456 ETH
Uptime: 98.50%
Status: active
```

### Tests unitaires
```bash
go test ./internal/kiln -v
```

## 📝 Logs

Le service produit des logs structurés :
```
2024/01/15 14:30:00 Configuration chargée avec succès!
2024/01/15 14:30:00 Base URL: https://api.kiln.fi
2024/01/15 14:30:00 Intervalle: 1m0s
2024/01/15 14:30:00 kiln client initialized with baseUrl : https://api.kiln.fi
2024/01/15 14:30:00 démarrage du collector intervalle : 1m0s
2024/01/15 14:30:00 compte à surveiller : [acc_demo]
2024/01/15 14:30:00 Première collecte
2024/01/15 14:30:00 📊 Collecte des données de 14:25:00 à 14:30:00
2024/01/15 14:30:00 Account: acc_demo | Status: active | Reward: 0.023456 ETH | Uptime: 98.50% | Window: 14:30:00
```

## 🔧 Développement

### Structure du projet
```
├── main.go                          # Point d'entrée
├── internal/
│   ├── config/                      # Gestion de la configuration
│   ├── kiln/                        # Client API Kiln
│   └── collector/                   # Collecteur de métriques
├── configs/
│   └── config.example.yaml         # Configuration exemple
└── README.md
```

### Ajout de nouvelles fonctionnalités

1. **Nouvelles métriques** : Étendre `RewardsSnapshot` dans `internal/kiln/client.go`
2. **Nouvelles règles** : Ajouter des champs dans `Rules` struct
3. **Nouveaux canaux d'alerte** : Étendre le système d'alertes

## 🚨 Alertes

### Configuration Discord
```yaml
alerts:
  kind: discord
  discord_webhook: "https://discord.com/api/webhooks/YOUR_WEBHOOK"
```

### Variables d'environnement pour la production
```bash
export KILN_API_TOKEN="production_token"
export DISCORD_WEBHOOK_URL="production_webhook"
export ALERT_KIND="discord"
```

## 🔒 Sécurité

- Les tokens API sont lus depuis les variables d'environnement
- Le fichier de configuration ne contient pas de données sensibles
- Support des timeouts HTTP configurables

## 📋 TODO / Roadmap

- [ ] Implémentation complète de l'API Kiln réelle
- [ ] Système d'alertes avancé
- [ ] Métriques Prometheus
- [ ] Dockerfile et docker-compose
- [ ] Tests d'intégration
- [ ] Documentation API

## 🤝 Contribution

1. Fork le projet
2. Créer une branche feature (`git checkout -b feature/AmazingFeature`)
3. Commit vos changements (`git commit -m 'Add some AmazingFeature'`)
4. Push vers la branche (`git push origin feature/AmazingFeature`)
5. Ouvrir une Pull Request

## 📞 Support

Pour toute question ou problème, ouvrez une issue sur le repository.

---

**Développé avec ❤️ en Go pour Kiln**