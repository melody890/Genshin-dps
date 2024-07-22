package event

type Event int

const (
	OnEnemyHit     Event = iota // target, AttackEvent
	OnPlayerHit                 // char, AttackEvent
	OnGadgetHit                 // target, AttackEvent
	OnEnemyDamage               // target, AttackEvent, amount, crit
	OnGadgetDamage              // target, AttackEvent
	OnApplyAttack               // AttackEvent
	// reaction related
	// OnReactionOccured // target, AttackEvent
	// OnTransReaction   // target, AttackEvent
	// OnAmpReaction     // target, AttackEvent

	OnAuraDurabilityAdded    // target, ele, durability
	OnAuraDurabilityDepleted // target, ele
	// OnReaction               // target, AttackEvent, ReactionType
	ReactionEventStartDelim
	OnOverload           // target, AttackEvent
	OnSuperconduct       // target, AttackEvent
	OnMelt               // target, AttackEvent
	OnVaporize           // target, AttackEvent
	OnFrozen             // target, AttackEvent
	OnElectroCharged     // target, AttackEvent
	OnSwirlHydro         // target, AttackEvent
	OnSwirlCryo          // target, AttackEvent
	OnSwirlElectro       // target, AttackEvent
	OnSwirlPyro          // target, AttackEvent
	OnCrystallizeHydro   // target, AttackEvent
	OnCrystallizeCryo    // target, AttackEvent
	OnCrystallizeElectro // target, AttackEvent
	OnCrystallizePyro    // target, AttackEvent
	OnAggravate          // target, AttackEvent
	OnSpread             // target, AttackEvent
	OnQuicken            // target, AttackEvent
	OnBloom              // target, AttackEvent
	OnHyperbloom         // target, AttackEvent
	OnBurgeon            // target, AttackEvent
	OnBurning            // target, AttackEvent
	OnShatter            // target, AttackEvent; at the end to simplify all reaction event subs since it's normally not considered as an elemental reaction
	ReactionEventEndDelim
	OnDendroCore // Gadget
	// other stuff
	OnStamUse          // abil
	OnShielded         // shield
	OnShieldBreak      // shield break
	OnConstructSpawned // nil
	OnCharacterSwap    // prev, next
	OnParticleReceived // particle
	OnEnergyChange     // character_received_index, pre_energy, energy_change, src (post-energy available in character_received), is_particle (boolean)
	OnTargetDied       // target, AttackEvent
	OnTargetMoved      // target
	OnCharacterHit     // nil <- this is for when the character is going to get hit but might be shielded from dmg
	OnCharacterHurt    // amount
	OnHPDebt           // target character, amount
	OnHeal             // src char, target character, amount, overheal, amount_before_debt
	OnPlayerPreHPDrain // Draininfo to modify
	OnPlayerHPDrain    // DrainInfo
	// ability use
	OnActionFailed // ActiveCharIndex, action.Action, param, action.ActionFailure
	OnActionExec   // ActiveCharIndex, action.Action, param
	OnSkill        // nil
	OnBurst        // nil
	OnAttack       // nil
	OnChargeAttack // nil
	OnPlunge       // nil
	OnAimShoot     // nil
	OnDash
	// sim stuff
	OnInitialize  // nil
	OnStateChange // prev, next
	OnEnemyAdded  // t
	OnTick
	OnSimEndedSuccessfully // nil
	EndEventTypes          // elim
)
