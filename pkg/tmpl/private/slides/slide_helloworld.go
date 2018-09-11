package slides

const (
	HelloWorld = `{{define "slide"}}
	<div class="slides">
            <section class="container" data-transition="slide-in fade-out">
                <div class="row">
                    <h1 class="col-12-12">Projet MEDTUNE</h1>
                    <h4 class="col-12-12">Présentation ICL Nancy</h4>
                    <h6 class="col-12-12 slide-1">12 Septembre 2018</h6>
                    <span id="adresse" class="col-6-12">Agence SII EST : 4
                        , Rue de Sarrelouis - 67000 Strasbourg</span>
                    <span id="coord" class="col-6-12"> www.groupe-sii.com </span>
                </div>
            </section>
            <!--slide 2 : Ordre du jour-->
            <section class="container" data-transition="convex">
                <div class="row">
                    <div class="col-3-12 barre barre--violette barre--active"></div>
                    <div class="col-3-12 barre barre--verte"></div>
                    <div class="col-3-12 barre barre--grise"></div>
                    <div class="col-3-12 barre barre--orange"></div>
                    <h3 class="col-12-12 barre--violette-text">Ordre du jour</h3>
                    <span class="col-12-12 slide-2">Constat de départ</span>
                    <span class="col-12-12 slide-2">Pourquoi ce projet ?</span>
                    <span class="col-12-12 slide-2">Aspects techniques</span>
                    <span class="col-12-12 slide-2">Description du projet</span>
                    <span class="col-12-12 slide-2">Démonstrations</span>
                </div>
            </section>
            <!--slide 3 : Constat de départ et objectif majeur-->
            <section class="container" data-transition="slide-out fade-in">
                <div class="row">
                    <div class="col-3-12 barre barre--violette barre--active"></div>
                    <div class="col-3-12 barre barre--verte"></div>
                    <div class="col-3-12 barre barre--grise"></div>
                    <div class="col-3-12 barre barre--orange"></div>
                    <h3 class="col-12-12 barre--violette-text">Constat de départ et objectif majeur</h3>
                    <span class="col-12-12 slide-3">Constat de départ
                        <li>Un manque d’outils avancés accessibles par les acteurs dans le domaine de la recherche
                            scientifique
                            et de la santé.
                        </li>
                        <li>Une domination des big pharma dans le domaine de la santé.
                        </li>
                    </span>
                    <span class="col-12-12 slide-3">Objectif majeur
                        <li>En offrant un service contenant un ensemble d’outils intelligents, permettant d’exploiter
                            plusieurs
                            types d’information ou de document.
                        </li>
                    </span>
                    <span class="col-12-12 slide-3">Première cible :
                        l’imagerie médicale
                    </span>
                </div>
            </section>
            <!--slide 4 : Pourquoi ce projet?-->
            <section class="container" data-transition="concave">
                <div class="row">
                    <div class="col-3-12 barre barre--violette barre--active"></div>
                    <div class="col-3-12 barre barre--verte"></div>
                    <div class="col-3-12 barre barre--grise"></div>
                    <div class="col-3-12 barre barre--orange"></div>
                    <h3 class="col-12-12 barre--violette-text">Pourquoi ce projet?</h3>
                    <img class="col-4-12 col-plus-1" src="/static/images/presentation/lesechos.png"></img>
                    <img class="col-4-12 col-plus-2 " src="/static/images/presentation/citationIA.png"></img>
                    <p class="col-12-12 slide-4">Grâce au <b>« machine learning assisté»</b>, […] l’algorithme sera
                        capable de
                        reconnaitre et d’interpréter des
                        clichés, mais aussi d’apprendre tout seul des « cas » en comparant une image aux milliers –
                        voire
                        aux millions – d’autres entrées dans sa base, et proposer un pré diagnostic. […]</p>
                    <p class="col-12-12 slide-4">Sur un point, tout le monde est à peu près d’accord : <b>l’IA sera un
                            outil d’aide à la décision</b> qui améliorera la prise en charge des patients.</p>
                    <p class="col-12-12 slide-4">[...] le nombre d’actes d’imagerie médicale en France, 80 millions
                        aujourd’hui,
                        <b>augmente chaque année
                            d’environ 4 %.</b></p>
                    <p class="col-12-12 slide-4">Les algorithmes doivent aider les praticiens à gérer cette inflation
                        d’examens
                        qu’ils ont de plus en
                        plus de difficultés à absorber.</p>
                    <p class="col-12-12 slide-4">Pour des tâches répétitives, le logiciel va augmenter la précision du
                        diagnostic. Il sera capable de
                        trier rapidement les images qui ne présentent aucune zone à problème, laissant plus de temps au
                        radiologue
                        pour se concentrer sur les cas complexes et la pathologie du patient.</p>
                    <p class="col-12-12 slide-4">[…] deux limites :<b>« un algorithme ne peut donner des renseignements
                            que sur ce
                            qu’il connait</b>. Par ailleurs,
                        s’il réalise un calcul à partir de données qui ne sont pas justes, les résultats seront
                        faussés.
                        »
                    </p>

                </div>
            </section>
            <!--slide 5 : Pourquoi ce projet?-->
            <section class="container" data-transition="zoom">
                <div class="row">
                    <div class="col-3-12 barre barre--violette barre--active"></div>
                    <div class="col-3-12 barre barre--verte"></div>
                    <div class="col-3-12 barre barre--grise"></div>
                    <div class="col-3-12 barre barre--orange"></div>
                    <h3 class="col-12-12 barre--violette-text">Pourquoi ce projet?</h3>
                    <span class="col-12-12">De nombreuses publications
                        scientifiques sur le sujet
                    </span>
                    <img class="col-5-12" src="/static/images/presentation/naturescience.png"></img>
                    <p class="col-5-12 col-plus-1 slide-4">Machine learning
                        <i>AI for medical imaging goes deep</i>
                        <i>Nature Medecine 24, 539-540(2018)</i>
                    </p>
                    <p class="col-12-12 slide-4">An artificial intelligence (AI) using a deep-learning
                        approach can classify retinal images from optical
                        coherence tomography for early diagnosis of retinal diseases and has the potential to be used
                        in
                        other image-based medical diagnoses.
                    </p>
                    <div class="col-2-12 col-plus-1">
                        <img src="/static/images/presentation/ibm.png"></img>
                        <img src="/static/images/presentation/watson.png"></img>
                    </div>
                    <div class="col-2-12 col-plus-2">
                        <img src="/static/images/presentation/googleai.png"></img>
                        <img src="/static/images/presentation/googlebrain.png"></img>

                    </div>
                    <div class="col-2-12 col-plus-2">
                        <img src="/static/images/presentation/siemens.png"></img>
                        <img src="/static/images/presentation/siemensinfo.png"></img>
                    </div>
                </div>
            </section>
            <!--slide 6 : DESCRIPTION DU PROJET-->
            <section class="container" data-transition="convex">
                <div class="row">
                    <div class="col-3-12 barre barre--violette"></div>
                    <div class="col-3-12 barre barre--verte barre--active"></div>
                    <div class="col-3-12 barre barre--grise"></div>
                    <div class="col-3-12 barre barre--orange"></div>
                    <h3 class="col-12-12 barre--verte-text">DESCRIPTION DU PROJET</h3>
                    <span class="col-6-12 slide-6">Le point de départ :
                        <p>Des thèses scientifiques publiques traitant d’IA et d’imagerie décrivent :</p>
                        <li>
                            <b>But des thèses</b> : faire une validation du modèle avec un taux de confiance « correct
                            »
                            mais
                            peu optimisé (recherche fondamentale, sponsors, quantité des publications vs qualité).
                        </li>
                        <li>
                            <b>Un sujet scientifique</b> : par exemple, les pathologies pulmonaires.
                        </li>
                        <li>
                            <b>Un algorithme</b> : description ± succincte du programme et de sa philosophie.
                        </li>
                        <li>
                            <b>Une base documentaire</b> : référence à d’autres thèses ou articles scientifiques.
                        </li>
                        <li>
                            <b>Une base de données (imagerie et/ou données)</b> : anonymisées et accessibles en libre
                            service
                            sur des sites officiels (NIH, univ. Stanford, …).
                        </li>
                    </span>
                    <span class="col-5-12 col-plus-1 slide-6">Travail réalisé :
                        <li>Lire un certain nombre de thèses en fonction des « thèmes » ciblés (> 150 thèses
                            identifiées).
                        </li>
                        <li>Identifier le potentiel du travail présenté dans chacune.
                        </li>
                        <li>Evaluer l’amélioration qui peut être apportée à l’algorithme
                        </li>
                        <li>Identifier les prérequis techniques et fonctionnels pour la reproductibilité des résultats
                            trouvés.
                        </li>
                    </span>
                </div>
            </section>
            <!--slide 7 : DESCRIPTION DU PROJET-->
            <section class="container" data-transition="slide-in fade-in">
                <div class="row">
                    <div class="col-3-12 barre barre--violette"></div>
                    <div class="col-3-12 barre barre--verte barre--active"></div>
                    <div class="col-3-12 barre barre--grise"></div>
                    <div class="col-3-12 barre barre--orange"></div>
                    <h3 class="col-12-12 barre--verte-text">DESCRIPTION DU PROJET
                    </h3>
                    <span class="col-12-12">Bases d’informations utilisées à
                        ce jour
                    </span>
                    <div class="col-6-12 slide-6">
                        <p>
                            <b>ChestX (NIH)</b>
                            <li>14 pathologies pulmonaires</li>
                            <li>40 GB data</li>
                            <li>112.000 images</li>
                            <li>Machine learning sur 100.000 images</li>
                            <li>Validation sur 12.000 images</li>
                        </p>
                        <img src="/static/images/presentation/chest1.png"></img>
                        <img src="/static/images/presentation/chest2.png"></img>
                    </div>
                    <div class="col-5-12 col-plus-1 slide-6">
                        <p>
                            <b>MURA (Univ Stanford)</b>
                            <li>Problèmes musculo-squelettiques</li>
                            <li>4 GB data</li>
                            <li>38.000 images</li>
                            <li>Machine learning sur 34.000 images</li>
                            <li>Validation sur 4.000 images</li>
                        </p>
                        <img src="/static/images/presentation/mura1.PNG"></img>
                        <img src="/static/images/presentation/mura2.PNG"></img>
                    </div>
                </div>
            </section>
            <!--slide 8 : DESCRIPTION DU PROJET-->
            <section class="container" data-transition="slide-out fade-out">
                <div class="row">
                    <div class="col-3-12 barre barre--violette"></div>
                    <div class="col-3-12 barre barre--verte barre--active"></div>
                    <div class="col-3-12 barre barre--grise"></div>
                    <div class="col-3-12 barre barre--orange"></div>
                    <h3 class="col-12-12 barre--verte-text">DESCRIPTION DU PROJET</h3>
                    <ul class="col-12-12 slide-7">La mise en œuvre:
                        <li>
                            Récupération des informations (imagerie et données associées ) utilisées par la thèse.
                        </li>
                        <li>
                            Construction de l’algorithme de départ.
                        </li>
                        <li>Reproductibilité des résultats trouvés dans la thèse (indice de confiance).
                            <ul>
                                <li>Apprentissage (machine learning) sur 90% des images/données.</li>
                                <li>Réalisation d’un nombre important d’itérations (50 à 100 sur l’ensemble des
                                    données).</li>
                                <li>Obtention d’un indice de confiance.</li>
                            </ul>
                        </li>
                        <li>Amélioration de l’algorithme pour augmenter l’indice de confiance.
                            <ul>
                                <li>Apprentissage (machine learning) sur les mêmes 90% d’images/données.</li>
                                <li>Réalisation d’un nombre important d’itérations (50 à 100).</li>
                                <li>Obtention d’un indice de confiance significativement supérieur au précédent.</li>
                            </ul>
                        </li>
                        <li>Validation des résultats via de nouvelles données.
                            <ul>
                                <li>Utilisation des 10% d’images non traitées par l’algorithme.</li>
                                <li>Analyse de la finesse des résultats.</li>
                            </ul>
                        </li>
                    </ul>
                </div>
            </section>
            <!--slide 9 : DESCRIPTION DU PROJET-->
            <section class="container" data-transition="slide-in fade-out">
                <div class="row">
                    <div class="col-3-12 barre barre--violette"></div>
                    <div class="col-3-12 barre barre--verte barre--active"></div>
                    <div class="col-3-12 barre barre--grise"></div>
                    <div class="col-3-12 barre barre--orange"></div>
                    <h3 class="col-12-12 barre--verte-text">DESCRIPTION DU PROJET
                    </h3>
                    <span class="col-12-12 slide-9">Exemples de courbes de
                        machine learning sur imagerie
                    </span>
                    <img class="col-5-12 slide-6" src="/static/images/presentation/appmauvais.PNG"></img>
                    <img class="col-5-12 col-plus-1 slide-6" src="/static/images/presentation/appbon.PNG"></img>
                </div>
            </section>
            <!--slide 10 : DESCRIPTION DU PROJET-->
            <section class="container" data-transition="slide-out fade-in">
                <div class="row">
                    <div class="col-3-12 barre barre--violette"></div>
                    <div class="col-3-12 barre barre--verte barre--active"></div>
                    <div class="col-3-12 barre barre--grise"></div>
                    <div class="col-3-12 barre barre--orange"></div>
                    <h3 class="col-12-12 barre--verte-text">DESCRIPTION DU PROJET
                    </h3>
                    <span class="col-12-12 slide-10">Organisation en mode
                        agile (Sprint*)
                        <p>Sprint : de la récupération des informations jusqu’à la livraison à l’utilisateur pour qu’il
                            teste
                            ses propres images sur la plateforme technique.
                        </p>
                        <i>* Le travail d’analyse des thèses de recherche est hors périmètre du sprint</i>
                    </span>
                    <img class="col-9-12 col-plus-1" src="/static/images/presentation/gantt.png"></img>
                </div>
            </section>
            <!--slide 11 : DESCRIPTION DU PROJET-->
            <section class="container" data-transition="zoom-in">
                <div class="row">
                    <div class="col-3-12 barre barre--violette"></div>
                    <div class="col-3-12 barre barre--verte barre--active"></div>
                    <div class="col-3-12 barre barre--grise"></div>
                    <div class="col-3-12 barre barre--orange"></div>
                    <h3 class="col-12-12 barre--verte-text">DESCRIPTION DU PROJET
                    </h3>
                    <span class="col-12-12">Informations sur un Sprint</span>
                </div>
                <table>
                    <caption>
                        <b>* X</b> dépend de la taille du modèle et des données</caption>
                    <tr class="slide-11">
                        <th>Etapes</th>
                        <th>Durée moy. (j)</th>
                        <th>Charge moy. (j)</th>
                        <th>Profils</th>
                    </tr>
                    <tr class="slide-11">
                        <td>
                            <i>Modéliser la théorie</i>
                        </td>
                        <td>2 - 3 semaines</td>
                        <td>12,5</td>
                        <td>Développeur Python/Go/C++ (junior à senior) Connaissance de Tensor Flow</td>
                    </tr>
                    <tr class="slide-11">
                        <td>
                            <i>Traiter les données brutes</i>
                        </td>
                        <td>2 à 3 semaines</td>
                        <td>13,5</td>
                        <td>Développeur avec compétences en traitement des données (middle à senior)</td>
                    </tr>
                    <tr class="slide-11">
                        <td>
                            <i>Entraîner le modèle</i>
                        </td>
                        <td>12h à 1 semaine et plus</td>
                        <td>0,5 + X*1,5h</td>
                        <td>Développeur Python/Go/C++ (middle à senior) Idéalement expérience de machine learning</td>
                    </tr>
                    <tr class="slide-11">
                        <td>
                            <i>Évaluer et valider le modèle</i>
                        </td>
                        <td>-</td>
                        <td>1 + X*1,5h</td>
                        <td>Développeur Python/Go/C++ (middle à senior) Idéalement expérience de machine learning</td>
                    </tr>
                    <tr class="slide-11">
                        <td>
                            <i>Figer le modèle</i>
                        </td>
                        <td>1 jour</td>
                        <td>0,5</td>
                        <td>Développeur Python/Go/C++ (junior à senior)</td>
                    </tr>
                    <tr class="slide-11">
                        <td>
                            <i>Délivrer le modèle en PROD</i>
                        </td>
                        <td>Début d'un nouveau jalon</td>
                        <td>0,5</td>
                        <td>Développeur Go/Python/C++ (junior à senior)</td>
                    </tr>
                </table>
                <span id=dureecharge>
                    <a>
                        <b>
                            <i>Durée totale</i>
                        </b> : 50 à 60 jours.</a>
                    <a>
                        <b>
                            <i>Charge totale</i>
                        </b> : 30 à 50 jours.</a>
                </span>
            </section>
            <!--slide 12: ASPECTS TECHNIQUES-->
            <section class="container" data-transition="convex">
                <div class="row">
                    <div class="col-3-12 barre barre--violette"></div>
                    <div class="col-3-12 barre barre--verte"></div>
                    <div class="col-3-12 barre barre--grise barre--active"></div>
                    <div class="col-3-12 barre barre--orange"></div>
                    <h3 class="col-12-12 barre--grise-text">ASPECTS TECHNIQUES
                    </h3>
                    <span class="col-12-12">Les Frameworks de Machine Learning</span>
                    <div class="col-12-12">
                        <table>
                            <caption>
                                <b>* X</b> dépend de la taille du modèle et des données</caption>
                            <tr class="slide-11">
                                <th>Rang</th>
                                <th>Famework</th>
                                <th>Entreprise</th>
                                <th>Popularité (GitHub)</th>
                                <th>Développement/Déployement</th>
                            </tr>
                            <tr class="slide-11">
                                <td>1</td>
                                <td>
                                    <i>TensorFlow</i>
                                </td>
                                <td>Google</td>
                                <td>109000 stars</td>
                                <td>Fonctions de développement avancées<br>Déployement intégré à l'outil</td>
                            </tr>
                            <tr class="slide-11">
                                <td>2</td>
                                    <td>
                                        <i>Scikit-learn </i>
                                    </td>
                                    <td>Contributeurs Open source</td>
                                    <td>30500 stars</td>
                                    <td>Fonctions de développement avancées<br>Déployement non intégré à l'outil</td>
                            </tr>
                            <tr class="slide-11">
                                    <td>3</td>
                                    <td>
                                        <i>Pytorch</i>
                                    </td>
                                    <td>Facebook</td>
                                    <td>18500 stars</td>
                                    <td>Développement pour la recherche<br>Déployement non intégré à l'outil</td>
                            </tr>
                            <tr class="slide-11">
                                <td>4</td>
                                <td>
                                    <i>CNTK</i>
                                </td>
                                <td>Microsoft</td>
                                <td>15100 stars</td>
                                <td>Fonctions de développement avancées<br>Déployement intégré à l'outil</td>
                                </tr>
                            <tr class="slide-11">
                                <td>5</td>
                                <td>
                                    <i>Theano</i>
                                </td>
                                <td>Université McGilles(Montréal)</td>
                                <td>8500 stars</td>
                                <td>Développement possible<br>Déployement non intégré à l'outil</td>
                            </tr>
                            
                            
                        </table>
                    </div>
                </div>
            </section>
            <!--slide 13: ASPECTS TECHNIQUES-->
            <section class="container" data-transition="concave">
                <div class="row">
                    <div class="col-3-12 barre barre--violette"></div>
                    <div class="col-3-12 barre barre--verte"></div>
                    <div class="col-3-12 barre barre--grise barre--active"></div>
                    <div class="col-3-12 barre barre--orange"></div>
                    <h3 class="col-12-12 barre--grise-text">ASPECTS TECHNIQUES
                    </h3>
                    <span class="col-12-12">TensorFlow</span>
                    <div class="col-2-12">
                        <img src="/static/images/presentation/tensorflow.jpg">
                    </div>
                    <div class="col-10-12">
                        <img src="/static/images/presentation/Tensorflow slide Medtune ICL NAncy.png">
                    </div>
                </div>
            </section>
            <!--slide 14: ASPECTS TECHNIQUES-->
            <section class="container" data-transition="zoom-in">
                <div class="row">
                    <div class="col-3-12 barre barre--violette"></div>
                    <div class="col-3-12 barre barre--verte"></div>
                    <div class="col-3-12 barre barre--grise barre--active"></div>
                    <div class="col-3-12 barre barre--orange"></div>
                    <h3 class="col-12-12 barre--grise-text">ASPECTS TECHNIQUES
                    </h3>
                    <div class="col-12-12">
                        <img id="img-slide-14" src="/static/images/presentation/entrainement-schema.PNG">
                    </div>
                </div>
            </section>
            <!--slide 15: ASPECTS TECHNIQUES-->
            <section class="container" data-transition="zoom-out">
                <div class="row">
                    <div class="col-3-12 barre barre--violette"></div>
                    <div class="col-3-12 barre barre--verte"></div>
                    <div class="col-3-12 barre barre--grise barre--active"></div>
                    <div class="col-3-12 barre barre--orange"></div>
                    <h3 class="col-12-12 barre--grise-text">ASPECTS TECHNIQUES
                    </h3>
                    <div class="col-12-12">
                        <img id="img-slide-12" src="/static/images/presentation/technos-deployment.png">
                    </div>
                </div>
            </section>
            <!--slide 16: ASPECTS TECHNIQUES-->
            <section class="container" data-transition="zoom-out">
                <div class="row">
                    <div class="col-3-12 barre barre--violette"></div>
                    <div class="col-3-12 barre barre--verte"></div>
                    <div class="col-3-12 barre barre--grise barre--active"></div>
                    <div class="col-3-12 barre barre--orange"></div>
                    <h3 class="col-12-12 barre--grise-text">ASPECTS TECHNIQUES
                    </h3>
                    <div class="col-12-12">
                        <img class="img-slide-12" src="/static/images/presentation/services.png">
                    </div>
                </div>
            </section>
            <!--slide 17: Passons à la démo-->
            <section class="container" data-transition="fade">
                <div class="row">
                    <div class="col-3-12 barre barre--violette"></div>
                    <div class="col-3-12 barre barre--verte"></div>
                    <div class="col-3-12 barre barre--grise"></div>
                    <div class="col-3-12 barre barre--orange barre--active"></div>
                    <h3 class="col-12-12 barre--orange-text">Passons à la démo
                    </h3>
                </div>
            </section>
        </div>
{{end}}`
)
