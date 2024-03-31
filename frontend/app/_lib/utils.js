
   // Exemple d'utilisation de fetch pour envoyer une requête au serveur
   export async function getSessionUser() {
    const token = localStorage.getItem('token');
    const response = await fetch('http://localhost:8080/auth/usersessions?token=' + token, {
       method: 'GET',
       headers: {
         'Content-Type': 'application/json',
       },
    });
    
    if (!response.ok) {
       throw new Error('Failed to fetch session');
    }
    
    const user = await response.json();
   //  console.log("user in session", user); // Afficher les informations de l'utilisateur
    return user; // Retourner les données de l'utilisateur
}