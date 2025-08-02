const fs = require('fs');

const networks = ['Commercial', 'Medicare', 'Tricare'];
const providerNetworks = [];

// Generate network data with realistic distributions
for (let i = 1; i <= 11921; i++) {
  const providerId = `P${String(i).padStart(4, '0')}`;
  
  // Read provider data to get county info
  const providers = JSON.parse(fs.readFileSync('data/providers.json', 'utf8'));
  const provider = providers.find(p => p.provider_id === providerId);
  
  if (!provider) continue;
  
  const county = provider.county;
  const specialty = provider.provider_type;
  
  // Different network participation rates based on county and specialty
  let commercialRate = 0.8; // 80% base rate
  let medicareRate = 0.6;   // 60% base rate  
  let tricareRate = 0.3;    // 30% base rate
  
  // Urban counties have higher commercial participation
  if (['Johnson', 'Sedgwick', 'Shawnee', 'Douglas'].includes(county)) {
    commercialRate = 0.95;
    medicareRate = 0.85;
    tricareRate = 0.6;
  }
  
  // Rural counties have lower commercial, higher medicare
  if (!['Johnson', 'Sedgwick', 'Shawnee', 'Douglas', 'Leavenworth', 'Reno', 'Saline'].includes(county)) {
    commercialRate = 0.5;
    medicareRate = 0.9;
    tricareRate = 0.2;
  }
  
  // Military areas (near bases) have higher Tricare
  if (['Geary', 'Riley', 'Leavenworth'].includes(county)) {
    tricareRate = 0.8;
  }
  
  // Specialists have different patterns
  if (['Cardiology', 'Neurology', 'Oncology'].includes(specialty)) {
    commercialRate *= 1.1;
    medicareRate *= 0.8;
  }
  
  // Generate network participation
  networks.forEach(network => {
    let participationRate;
    if (network === 'Commercial') participationRate = commercialRate;
    else if (network === 'Medicare') participationRate = medicareRate;
    else participationRate = tricareRate;
    
    if (Math.random() < participationRate) {
      const isTerminated = Math.random() < 0.08; // 8% terminated
      
      let effectiveDate, terminationDate, terminationReason;
      
      if (isTerminated) {
        const termYear = 2020 + Math.floor(Math.random() * 3);
        const termMonth = Math.floor(Math.random() * 12) + 1;
        const termDay = Math.floor(Math.random() * 28) + 1;
        
        effectiveDate = `${termYear - 1}-01-01T00:00:00Z`;
        terminationDate = `${termYear}-${String(termMonth).padStart(2, '0')}-${String(termDay).padStart(2, '0')}T00:00:00Z`;
        terminationReason = 'Left Network';
      } else {
        effectiveDate = '2023-01-01T00:00:00Z';
        terminationDate = '9999-12-31T00:00:00Z';
        terminationReason = '';
      }
      
      providerNetworks.push({
        provider_id: providerId,
        network_id: network,
        effective_date: effectiveDate,
        termination_date: terminationDate,
        termination_reason: terminationReason
      });
    }
  });
}

fs.writeFileSync('data/provider_networks.json', JSON.stringify(providerNetworks, null, 2));
console.log(`Generated ${providerNetworks.length} realistic network relationships`);