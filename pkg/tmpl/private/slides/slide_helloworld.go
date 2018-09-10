package slides

const (
	HelloWorld = `{{define "slide"}}
	<div class="slides">
            <section data-transition="slide-in fade-out">
                <div class="mdl-grid">
                    <h1 class="mdl-cell mdl-cell--12-col mdl-cell--top">Projet MEDTUNE</h1>
                    <h4 class="mdl-cell mdl-cell--12-col mdl-cell--top">Présentation ICL Nancy</h4>
                    <h6 class="mdl-cell mdl-cell--12-col mdl-cell--top">12 Septembre 2018</h6>
                    <span id="adresse" class="mdl-cell mdl-cell--6-col mdl-color-text--blue-900">Agence SII EST : 4
                        ,rue de Sarrelouis - 67000 Strasbourg</span>
                    <span id="coord" class="mdl-cell mdl-cell--6-col mdl-color-text--blue-900">Tél : 03 90 23 62 62 Fax
                        : 03 88 32 07 66 www.groupe-sii.com </span>
                </div>
            </section>
            <!--slide 2 : Ordre du jour-->
            <section data-transition="convex">
                <div class="mdl-grid">
                    <h3 class="mdl-cell mdl-cell--12-col mdl-cell--top slide-title">Ordre du jour</h3>
                    <span class="mdl-cell mdl-cell--6-col mdl-color-text--blue-900 slide-2">Pourquoi ce projet ?</span>
                    <span class="mdl-cell mdl-cell--6-col mdl-color-text--blue-900 slide-2">Aspects techniques</span>
                    <span class="mdl-cell mdl-cell--6-col mdl-color-text--blue-900 slide-2">Description du projet</span>
                    <span class="mdl-cell mdl-cell--6-col mdl-color-text--blue-900 slide-2">Démonstrations</span>
                </div>
            </section>
            <!--slide 3 : Constat de départ et objectif majeur-->
            <section data-transition="convex">
                <div class="mdl-grid">
                    <h3 class="mdl-cell mdl-cell--12-col mdl-cell--top">Constat de départ et objectif majeur</h3>
                    <span class="mdl-cell mdl-cell--12-col mdl-color-text--blue-900 slide-3">Constat de départ
                        <li>Un manque d’outils avancés accessibles par les acteurs dans le domaine de la recherche
                            scientifique
                            et de la santé.
                        </li>
                        <li>Une domination des big pharma dans le domaine de la santé.
                        </li>
                    </span>
                    <span class="mdl-cell mdl-cell--12-col mdl-color-text--blue-900 slide-3">Objectif majeur
                        <li>En offrant un service contenant un ensemble d’outils intelligents, permettant d’exploiter
                            plusieurs
                            types d’information ou de document.
                        </li>
                    </span>
                    <span class="mdl-cell mdl-cell--12-col mdl-color-text--blue-900 slide-3">Première cible :
                        l’imagerie médicale
                    </span>
                </div>
            </section>
            <!--slide 4 : Pourquoi ce projet?-->
            <section data-transition="convex">
                <div class="mdl-grid">
                    <h3 class="mdl-cell mdl-cell--12-col mdl-cell--top slide-title">Pourquoi de projet?</h3>
                    <img class="mdl-cell mdl-cell--4-col mdl-cell--4-offset slide-4-img" src="../static/images/presentation/lesechos.png"></img>
                    <img class="mdl-cell mdl-cell--4-col" src="../static/images/presentation/citationIA.png"></img>
                    <p class="slide-4">Grâce au « machine learning assisté», […] l’algorithme sera capable de
                        reconnaitre et d’interpréter des
                        clichés, mais aussi d’apprendre tout seul des « cas » en comparant une image aux milliers –
                        voire
                        aux millions – d’autres entrées dans sa base, et proposer un pré diagnostic. […]</p>
                    <p class="slide-4">Sur un point, tout le monde est à peu près d’accord : l’IA sera un outil d’aide
                        à la décision qui améliorera
                        la prise en charge des patients.</p>
                    <p class="slide-4">[...] le nombre d’actes d’imagerie médicale en France, 80 millions aujourd’hui,
                        augmente chaque année
                        d’environ 4 %.</p>
                    <p class="slide-4">Les algorithmes doivent aider les praticiens à gérer cette inflation d’examens
                        qu’ils ont de plus en
                        plus de difficultés à absorber.</p>
                    <p class="slide-4">Pour des tâches répétitives, le logiciel va augmenter la précision du
                        diagnostic. Il sera capable de
                        trier rapidement les images qui ne présentent aucune zone à problème, laissant plus de temps au
                        radiologue
                        pour se concentrer sur les cas complexes et la pathologie du patient.</p>
                    <p class="slide-4">[…] deux limites : « un algorithme ne peut donner des renseignements que sur ce
                        qu’il connait. Par ailleurs,
                        s’il réalise un calcul à partir de données qui ne sont pas justes, les résultats seront
                        faussés.
                        »
                    </p>

                </div>
            </section>
            <!--slide 5 : Pourquoi ce projet?-->
            <section data-transition="convex">
                <div class="mdl-grid">
                    <h3 class="mdl-cell mdl-cell--12-col mdl-cell--top slide-title">Pourquoi de projet?</h3>
                    <span class="mdl-cell mdl-cell--12-col mdl-color-text--blue-900">De nombreuses publications
                        scientifiques sur le sujet
                    </span>
                    <img class="mdl-cell mdl-cell--4-col" src="../static/images/presentation/naturescience.png"></img>
                    <p class="mdl-cell mdl-cell--8-col slide-4">Machine learning
                        <i>AI for medical imaging goes deep</i>
                        <i>Nature Medecine 24, 539-540(2018)</i>
                    </p>
                    <p class="mdl-cell mdl-cell--12-col slide-4">An artificial intelligence (AI) using a deep-learning
                        approach can classify retinal images from optical
                        coherence tomography for early diagnosis of retinal diseases and has the potential to be used
                        in
                        other image-based medical diagnoses.
                    </p>
                </div>
                <div class="slide-5">
                    <img class="slide-5-img-2" src="../static/images/presentation/ibm.png"></img>
                    <img class="slide-5-img-2" src="../static/images/presentation/googleai.png"></img>
                    <img class="slide-5-img-2" src="../static/images/presentation/siemens.png"></img>
                </div>
                <div class="slide-5">
                    <img class="slide-5-img-2" src="../static/images/presentation/watson.png"></img>
                    <img class="slide-5-img-2" src="../static/images/presentation/googlebrain.png"></img>
                    <img class="slide-5-img-2" src="../static/images/presentation/siemensinfo.png"></img>
                </div>
            </section>
            <!--slide 6 : DESCRIPTION DU PROJET-->
            <section>
                <div class="mdl-grid">
                    <h3 class="mdl-cell mdl-cell--12-col mdl-cell--top">DESCRIPTION DU PROJET</h3>
                    <span class="mdl-cell mdl-cell--6-col mdl-color-text--blue-900 slide-6">Le point de départ:
                        <p>Des thèses scientifiques publiques traitant d’IA et d’imagerie décrivent :</p>
                        <li>
                            <b>But des thèses</b>: faire une validation du modèle avec un taux de confiance « correct »
                            mais
                            peu optimisé (recherche fondamentale, sponsors, quantité des publications vs qualité).
                        </li>
                        <li>
                            <b>Un sujet scientifique</b>: par exemple, les pathologies pulmonaires.
                        </li>
                        <li>
                            <b>Un algorithme</b>: description ± succincte du programme et de sa philosophie.
                        </li>
                        <li>
                            <b>Une base documentaire</b>: référence à d’autres thèses ou articles scientifiques.
                        </li>
                        <li>
                            <b>Une base de données (imagerie et/ou données)</b>: anonymisées et accessibles en libre
                            service
                            sur des sites officiels (NIH, univ. Stanford, …).
                        </li>
                    </span>
                    <span class="mdl-cell mdl-cell--6-col mdl-color-text--blue-900 slide-6">Travail réalisé :
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
            <section>
                <div class="mdl-grid">
                    <h3 class="mdl-cell mdl-cell--12-col mdl-cell--top slide-title">DESCRIPTION DU PROJET
                    </h3>
                    <span class="mdl-cell mdl-cell--12-col mdl-color-text--blue-900">Bases d’informations utilisées à
                        ce jour
                    </span>
                    <div class="mdl-cell mdl-cell--6-col slide-6">
                        <p>
                            <b>ChestX (NIH)</b>
                            <li>14 pathologies pulmonaires</li>
                            <li>40 GB data</li>
                            <li>112.000 images</li>
                            <li>Machine learning sur 100.000 images</li>
                            <li>Validation sur 12.000 images</li>
                        </p>
                        <img src="../static/images/presentation/chest1.png"></img>
                        <img src="../static/images/presentation/chest2.png"></img>
                    </div>
                    <div class="mdl-cell mdl-cell--6-col slide-6">
                        <p>
                            <b>MURA (Univ Stanford)</b>
                            <li>Problèmes musculo-squelettiques</li>
                            <li>4 GB data</li>
                            <li>38.000 images</li>
                            <li>Machine learning sur 34.000 images</li>
                            <li>Validation sur 4.000 images</li>
                        </p>
                        <img src="../static/images/presentation/mura1.PNG"></img>
                        <img src="../static/images/presentation/mura2.PNG"></img>
                    </div>
                </div>
            </section>
            <!--slide 8 : DESCRIPTION DU PROJET-->
            <section>
                <div class="mdl-grid">
                    <h3 class="mdl-cell mdl-cell--12-col mdl-cell--top">DESCRIPTION DU PROJET</h3>
                    <ul class="mdl-cell mdl-cell--12-col mdl-color-text--blue-900 slide-7">La mise en œuvre:
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
            <section>
                <div class="mdl-grid">
                    <h3 class="mdl-cell mdl-cell--12-col mdl-cell--top slide-title">DESCRIPTION DU PROJET
                    </h3>
                    <span class="mdl-cell mdl-cell--12-col mdl-color-text--blue-900 slide-9">Exemples de courbes de
                        machine learning sur imagerie
                    </span>
                    <img class="mdl-cell mdl-cell--6-col slide-6" src="../static/images/presentation/appmauvais.PNG"></img>
                    <img class="mdl-cell mdl-cell--6-col slide-6" src="../static/images/presentation/appbon.PNG"></img>
                </div>
            </section>
            <!--slide 10 : DESCRIPTION DU PROJET-->
            <section>
                <div class="mdl-grid">
                    <h3 class="mdl-cell mdl-cell--12-col mdl-cell--top slide-title">DESCRIPTION DU PROJET
                    </h3>
                    <span class="mdl-cell mdl-cell--12-col mdl-color-text--blue-900 slide-10">Organisation en mode
                        agile (Sprint*)
                        <p>Sprint : de la récupération des informations jusqu’à la livraison à l’utilisateur pour qu’il
                            teste
                            ses propres images sur la plateforme technique.
                        </p>
                        <i>* Le travail d’analyse des thèses de recherche est hors périmètre du sprint</i>
                    </span>
                    <span class="mdl-cell mdl-cell--2-col"></span>
                    <img class="mdl-cell mdl-cell--9-col mdl-cell--top" src="../static/images/presentation/gantt.png"></img>
                </div>
            </section>
            <!--slide 11 : DESCRIPTION DU PROJET-->
            <section>
                <div class="mdl-grid">
                    <h3 class="mdl-cell mdl-cell--12-col mdl-cell--top slide-title">DESCRIPTION DU PROJET
                    </h3>
                    <span class="mdl-cell mdl-cell--12-col mdl-color-text--blue-900">Informations sur un Sprint</span>
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
                        <td></td>
                        <td>12,5</td>
                        <td>Développeur Python/Go/C++ (junior à senior) Connaissance de Tensor Flow</td>
                    </tr>
                    <tr class="slide-11">
                        <td>
                            <i>Traiter les données brutes</i>
                        </td>
                        <td></td>
                        <td>13,5</td>
                        <td>Développeur avec compétences en traitement des données (middle à senior)</td>
                    </tr>
                    <tr class="slide-11">
                        <td>
                            <i>Entraîner le modèle</i>
                        </td>
                        <td></td>
                        <td>0,5 + X*1,5h</td>
                        <td>Développeur Python/Go/C++ (middle à senior) Idéalement expérience de machine learning</td>
                    </tr>
                    <tr class="slide-11">
                        <td>
                            <i>Évaluer et valider le modèle</i>
                        </td>
                        <td></td>
                        <td>1 + X*1,5h</td>
                        <td>Développeur Python/Go/C++ (middle à senior) Idéalement expérience de machine learning</td>
                    </tr>
                    <tr class="slide-11">
                        <td>
                            <i>Figer le modèle</i>
                        </td>
                        <td></td>
                        <td>0,5</td>
                        <td>Développeur Python/Go/C++ (junior à senior)</td>
                    </tr>
                    <tr class="slide-11">
                        <td>
                            <i>Délivrer le modèle en PROD</i>
                        </td>
                        <td></td>
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
            <section>
                <div class="mdl-grid">
                    <h3 class="mdl-cell mdl-cell--12-col mdl-cell--top">ASPECTS TECHNIQUES
                    </h3>
                    <div class="mdl-cell mdl-cell--12-col">
                        <img id="img-slide-12" src="../static/images/presentation/beta-platform-diagram.png">
                    </div>
                </div>
            </section>
            <!--slide 13: ASPECTS TECHNIQUES-->
            <section>
                <div class="mdl-grid">
                    <h3 class="mdl-cell mdl-cell--12-col mdl-cell--top">ASPECTS TECHNIQUES
                    </h3>
                    <span class="mdl-cell mdl-cell--12-col mdl-color-text--blue-900">Les Frameworks Machine Learning</span>
                    <div class="mdl-cell mdl-cell--12-col">
                        <img src="../static/images/presentation/theano.jpg">
                        <img src="../static/images/presentation/cntk.png">
                        <img src="../static/images/presentation/caffe.png">
                        <img src="../static/images/presentation/pytorch.png">
                        <img src="../static/images/presentation/tensorflow.jpg">
                    </div>
                </div>
            </section>
            <!--slide 14: ASPECTS TECHNIQUES-->
            <section>
                <div class="mdl-grid">
                    <h3 class="mdl-cell mdl-cell--12-col mdl-cell--top">ASPECTS TECHNIQUES
                    </h3>
                    <span class="mdl-cell mdl-cell--12-col mdl-color-text--blue-900">TensorFlow</span>
                    <div class="mdl-cell mdl-cell--2-col mdl-cell--middle">
                        <img src="../static/images/presentation/tensorflow.jpg">
                    </div>
                    <div class="mdl-cell mdl-cell--10-col">
                        <img src="../static/images/presentation/Tensorflow slide Medtune ICL NAncy.png">
                    </div>
                </div>
            </section>
            <!--slide 15: ASPECTS TECHNIQUES-->
            <section>
                <div class="mdl-grid">
                    <h3 class="mdl-cell mdl-cell--12-col mdl-cell--top">ASPECTS TECHNIQUES
                    </h3>
                    <span class="mdl-cell mdl-cell--12-col mdl-color-text--blue-900">TensorFlow</span>
                    <div class="mdl-cell mdl-cell--2-col mdl-cell--middle">
                        <img src="../static/images/presentation/tensorflow.jpg">
                    </div>
                    <div class="mdl-cell mdl-cell--10-col">
                        <img src="../static/images/presentation/Tensorflow slide Medtune ICL NAncy.png">
                    </div>
                </div>
            </section>
        </div>
{{end}}`
)
